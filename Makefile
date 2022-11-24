DOCKER_NAME ?= variant64
DOCKER_APP_NAME_BE ?= server
DOCKER_APP_NAME_FE ?= web
DOCKER_APP_NAME_PS ?= proxy-server
DOCKER_APP_CONTAINER_BE = $(docker ps -a -q --filter ancestor=${DOCKER_NAME}/${DOCKER_APP_NAME_FE} --format="{{.ID}}")
FRONTEND_PKG_DIR_PATH = "frontend/packages"

# Deploy variables
ECR_REPO_ARN ?= 949672515001.dkr.ecr.us-east-2.amazonaws.com
ECS_CLUSTER ?= variant64
ECS_REGION ?= us-east-2
ECS_SERVICE_SERVER = server
ECS_SERVICE_WEB = web

all: build

# Help Command
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

# Setup Related Commands
install: # Install all dependencies for project.
	pipenv install -r requirements.txt

	(cd ./frontend/ && npm run bootstrap)

# Build Related Commands
build: build-server build-proxy build-client # Build all component's images.

build-server: # Build the server image.
	docker build -t ${DOCKER_NAME}/${DOCKER_APP_NAME_BE} ${DOCKER_APP_NAME_BE}/.

build-proxy: # Build the proxy image.
	docker build -t ${DOCKER_NAME}/${DOCKER_APP_NAME_PS} ${FRONTEND_PKG_DIR_PATH}/${DOCKER_APP_NAME_PS}/.

build-client: # Build the client image.
	docker build -t ${DOCKER_NAME}/${DOCKER_APP_NAME_FE} ${FRONTEND_PKG_DIR_PATH}/${DOCKER_APP_NAME_FE}/.

# Run Related Comamnds
run: run-server run-proxy run-client # Run all the component's containers.

run-server: stop-server build-server # Build and run the server.
	docker run -p 8000:8000 -d ${DOCKER_NAME}/${DOCKER_APP_NAME_BE}

run-proxy: stop-proxy build-proxy # Build and run the proxy.
	docker run -p 8001:8001 -d ${DOCKER_NAME}/${DOCKER_APP_NAME_PS}

run-client: stop-client build-client # Build and run the client.
	docker run -p 3000:3000 -d ${DOCKER_NAME}/${DOCKER_APP_NAME_FE}

# Stop Related Commands
stop: stop-server stop-proxy stop-client # Stop all the component's containers.

stop-server: # Stop the server container.
	docker rm $$(docker stop $$(docker ps -aq --filter ancestor=${DOCKER_NAME}/${DOCKER_APP_NAME_FE})) || true

stop-proxy: # Stop the proxy container.
	docker rm $$(docker stop $$(docker ps -aq --filter ancestor=${DOCKER_NAME}/${DOCKER_APP_NAME_PS})) || true

stop-client: # Stop the client container.
	docker rm $$(docker stop $$(docker ps -aq --filter ancestor=${DOCKER_NAME}/${DOCKER_APP_NAME_BE})) || true

# Test Related Comands
test: # Run all tests.
	cd server/ && go test ./...

# Deployment Related Commands
deploy-login:
	aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin ${ECR_REPO_ARN}

deploy-tag:
	$(eval REV=$(shell git rev-parse HEAD | cut -c1-7))
	docker tag ${DOCKER_NAME}/${DOCKER_APP_NAME_BE}:latest ${ECR_REPO_ARN}/${DOCKER_NAME}/${DOCKER_APP_NAME_BE}:${REV}
	docker tag ${DOCKER_NAME}/${DOCKER_APP_NAME_FE}:latest ${ECR_REPO_ARN}/${DOCKER_NAME}/${DOCKER_APP_NAME_FE}:${REV}
	docker tag ${DOCKER_NAME}/${DOCKER_APP_NAME_BE}:latest ${ECR_REPO_ARN}/${DOCKER_NAME}/${DOCKER_APP_NAME_BE}:latest
	docker tag ${DOCKER_NAME}/${DOCKER_APP_NAME_FE}:latest ${ECR_REPO_ARN}/${DOCKER_NAME}/${DOCKER_APP_NAME_FE}:latest

deploy-push-image:
	$(eval REV=$(shell git rev-parse HEAD | cut -c1-7))
	docker push ${ECR_REPO_ARN}/${DOCKER_NAME}/${DOCKER_APP_NAME_BE}:${REV}
	docker push ${ECR_REPO_ARN}/${DOCKER_NAME}/${DOCKER_APP_NAME_FE}:${REV}
	docker push ${ECR_REPO_ARN}/${DOCKER_NAME}/${DOCKER_APP_NAME_BE}:latest
	docker push ${ECR_REPO_ARN}/${DOCKER_NAME}/${DOCKER_APP_NAME_FE}:latest

deploy-refresh-task:
	aws ecs update-service --region $(ECS_REGION) --cluster ${ECS_CLUSTER} --service $(ECS_SERVICE_SERVER) --force-new-deployment
	aws ecs update-service --region $(ECS_REGION) --cluster ${ECS_CLUSTER} --service $(ECS_SERVICE_WEB) --force-new-deployment

deploy: build deploy-login deploy-tag deploy-push-image deploy-refresh-task # Deploy all the components.

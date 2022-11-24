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

install:
	pipenv install -r requirements.txt

	(cd ./frontend/ && npm run bootstrap)

build:
	docker build -t ${DOCKER_NAME}/${DOCKER_APP_NAME_BE} ${DOCKER_APP_NAME_BE}/.
	docker build -t ${DOCKER_NAME}/${DOCKER_APP_NAME_FE} ${FRONTEND_PKG_DIR_PATH}/${DOCKER_APP_NAME_FE}/.

run: stop build
	docker run -p 8000:8000 -d ${DOCKER_NAME}/${DOCKER_APP_NAME_BE}
	docker run -p 3000:3000 -d ${DOCKER_NAME}/${DOCKER_APP_NAME_FE}

stop:
	docker rm $$(docker stop $$(docker ps -aq --filter ancestor=${DOCKER_NAME}/${DOCKER_APP_NAME_FE})) || true
	docker rm $$(docker stop $$(docker ps -aq --filter ancestor=${DOCKER_NAME}/${DOCKER_APP_NAME_BE})) || true

test:
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

deploy: build deploy-login deploy-tag deploy-push-image deploy-refresh-task

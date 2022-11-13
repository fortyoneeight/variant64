DOCKER_NAME ?= variant64
DOCKER_APP_NAME_BE ?= server
DOCKER_APP_NAME_FE ?= web
DOCKER_APP_NAME_PS ?= proxy-server
DOCKER_APP_CONTAINER_BE = $(docker ps -a -q --filter ancestor=${DOCKER_NAME}/${DOCKER_APP_NAME_FE} --format="{{.ID}}")
FRONTEND_PKG_DIR_PATH = "frontend/packages"

ECR_REPO_ARN ?= 949672515001.dkr.ecr.us-east-2.amazonaws.com
ECS_CLUSTER ?= variant64
ECS_REGION ?= us-east-2
ECS_SERVICE_SERVER = server

all: build

build:
	docker build -t ${DOCKER_NAME}/${DOCKER_APP_NAME_BE} ${DOCKER_APP_NAME_BE}/.
	# docker build -t ${DOCKER_NAME}/${DOCKER_APP_NAME_PS} ${FRONTEND_PKG_DIR_PATH}/${DOCKER_APP_NAME_PS}/.
	docker build -t ${DOCKER_NAME}/${DOCKER_APP_NAME_FE} ${FRONTEND_PKG_DIR_PATH}/${DOCKER_APP_NAME_FE}/.

run: stop build
	docker run -p 8000:8000 -d ${DOCKER_NAME}/${DOCKER_APP_NAME_BE}
	# docker run -p 8001:8001 -d ${DOCKER_NAME}/${DOCKER_APP_NAME_PS}
	docker run -p 8100:8100 -d ${DOCKER_NAME}/${DOCKER_APP_NAME_FE}

stop:
	docker rm $$(docker stop $$(docker ps -aq --filter ancestor=${DOCKER_NAME}/${DOCKER_APP_NAME_FE})) || true
	# docker rm $$(docker stop $$(docker ps -aq --filter ancestor=${DOCKER_NAME}/${DOCKER_APP_NAME_PS})) || true
	docker rm $$(docker stop $$(docker ps -aq --filter ancestor=${DOCKER_NAME}/${DOCKER_APP_NAME_BE})) || true

test:
	cd server/ && go test ./...

# Deployment Related Commands
deploy-login:
	aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin ${ECR_REPO_ARN}

deploy-tag:
	$(eval REV=$(shell git rev-parse HEAD | cut -c1-7))
	docker tag ${DOCKER_NAME}/${DOCKER_APP_NAME_BE}:latest ${ECR_REPO_ARN}/${DOCKER_NAME}/${DOCKER_APP_NAME_BE}:latest

deploy-push-image:
	$(eval REV=$(shell git rev-parse HEAD | cut -c1-7))
	docker push ${ECR_REPO_ARN}/${DOCKER_NAME}/${DOCKER_APP_NAME_BE}:latest
	$(info -------------------------------)
	$(info -- Succesfully pushed image! --)
	$(info -------------------------------)

deploy-refresh-task:
	aws ecs update-service --region $(ECS_REGION) --cluster ${ECS_CLUSTER} --service $(ECS_SERVICE_SERVER) --force-new-deployment
	$(info ---------------------------)
	$(info -- Succesfully deployed! --)
	$(info ---------------------------)

deploy: build deploy-login deploy-tag deploy-push-image deploy-refresh-task

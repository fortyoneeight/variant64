DOCKER_NAME ?= variant64
DOCKER_APP_NAME_BE ?= server
DOCKER_APP_NAME_FE ?= web
DOCKER_APP_NAME_PS ?= proxy-server
DOCKER_APP_CONTAINER_BE = $(docker ps -a -q --filter ancestor=${DOCKER_NAME}/${DOCKER_APP_NAME_FE} --format="{{.ID}}")
FRONTEND_PKG_DIR_PATH = "frontend/packages"

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

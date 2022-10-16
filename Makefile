DOCKER_NAME ?= variant64
DOCKER_APP_NAME_BE ?= server
DOCKER_APP_NAME_FE ?= client
DOCKER_APP_CONTAINER_BE = $(docker ps -a -q --filter ancestor=${DOCKER_NAME}/${DOCKER_APP_NAME_FE} --format="{{.ID}}")

all: build

build:
	docker build -t ${DOCKER_NAME}/${DOCKER_APP_NAME_BE} server/.
	docker build -t ${DOCKER_NAME}/${DOCKER_APP_NAME_FE} client/.

run: build
	docker run -p 8000:8000 -d ${DOCKER_NAME}/${DOCKER_APP_NAME_BE}
	docker run -p 8100:8100 -d ${DOCKER_NAME}/${DOCKER_APP_NAME_FE}

stop:
	docker rm $$(docker stop $$(docker ps -aq --filter ancestor=${DOCKER_NAME}/${DOCKER_APP_NAME_FE}))
	docker rm $$(docker stop $$(docker ps -aq --filter ancestor=${DOCKER_NAME}/${DOCKER_APP_NAME_BE}))

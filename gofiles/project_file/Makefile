.PHONY: all build restart help init

# Variables
DIR=${shell pwd}
PATH_ENV=${realpath ${DIR}/app.env}
DOCKER_COMPOSE_PATH=${realpath ${DIR}/../../docker-compose.yml}

include ${PATH_ENV}
export

OS := $(shell uname -s)

ifeq ($(OS), Darwin)
    # Comandos para macOS
    SED_INPLACE := sed -i ''
else ifeq ($(OS), Linux)
    # Comandos para Linux / Ubuntu
    SED_INPLACE := sed -i
endif

# Colors
RED    := ${shell tput setaf 1}
GREEN  := ${shell tput setaf 2}
YELLOW := ${shell tput setaf 3}
BLUE   := ${shell tput setaf 4}
RESET  := ${shell tput sgr0}

help: ## Show this help
	@echo "${YELLOW}Usage:${RESET}"
	@echo "    make [build|restart|init]"

test:
	@echo ${SED_INPLACE}
	@echo ${DOCKER_COMPOSE_PATH}
	@${SED_INPLACE} 's/\$$\$$SERVICE_NAME\$$\$$/${SERVICE_NAME}/g' "${DOCKER_COMPOSE_PATH}"

restart:
	@echo "Restart service ${RED}${SERVICE_NAME}${RESET}"
	@echo "Stop service ${RED}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} stop ${SERVICE_NAME}
	@echo "Creating executeble ${BLUE}${SERVICE_NAME}${RESET}"
	@echo "Start service ${BLUE}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} --env-file ${PATH_ENV} start ${SERVICE_NAME}

init:
	@echo "Init service ${BLUE}${SERVICE_NAME}${RESET}"
	docker compose -f ${DOCKER_COMPOSE_PATH} --env-file ${PATH_ENV} up -d ${SERVICE_NAME}

build:
	@${SED_INPLACE} 's/\$$\$$SERVICE_NAME\$$\$$/${SERVICE_NAME}/g' "${DOCKER_COMPOSE_PATH}"
	@echo "Building ${GREEN}${SERVICE_NAME}${RESET}"
	@if [ ${CREATE_NETWORK} == true ]; then \
		echo "Creating network of service network name: ${GREEN}${SERVICE_NAME}${RESET}"; \
		docker network create --driver bridge \
			--subnet "${SUBNET_CONTAINER}" \
			--gateway "${GATEWAY_CONTAINER}" \
			"${DOCKER_EXTERNAL_NETWORK}"; \
	fi
	@echo "Building image service: ${GREEN}${SERVICE_NAME}${RESET}"
	docker compose -f ${DOCKER_COMPOSE_PATH} --env-file ${PATH_ENV} up --build ${SERVICE_NAME} 
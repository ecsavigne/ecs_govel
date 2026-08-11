.PHONY: all build restart help init restart_test

# Variables
DIR=${shell pwd}
PATH_ENV=${realpath ${DIR}/app.env}
DIR_PRINCIPAL=${realpath ${DIR}/../../}
DOCKER_COMPOSE_PATH=${realpath ${DIR_PRINCIPAL}/docker-compose.yml}
SUPERVISOR_CONFIG_PATH=${DIR_PRINCIPAL}/${SERVICE_NAME}.conf

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

NameServiceSuperVisor=name_service.conf
ReplaceNameServiceSuperVisor=${SERVICE_NAME}.conf
RENAME_SUPERVISOR_CONFIG=mv ${DIR_PRINCIPAL}/name_service.conf ${SUPERVISOR_CONFIG_PATH} || true

InSuperVisorConfig=<nameservice_APP_NAME_in_app.env><port_PORTS_in_app.env>
ReplaceInSuperVisorConfig=${SERVICE_NAME}${PORTS}

# Colors
RED    := ${shell tput setaf 1}
GREEN  := ${shell tput setaf 2}
YELLOW := ${shell tput setaf 3}
BLUE   := ${shell tput setaf 4}
RESET  := ${shell tput sgr0}

name_network := $(shell docker network ls | grep ${DOCKER_EXTERNAL_NETWORK} | awk '{print $$2}' )

help: ## Show this help
	@echo "${YELLOW}Usage:${RESET}"
	@echo "    make [build|restart|restart_test||]"
	@echo "    make build			\t- Create docker image for service ${SERVICE_NAME} and run service whith -d"
	@echo "    make restart			- Restart service ${SERVICE_NAME} and run service whith -d"
	@echo "    make restart_test	\t\t- Restart service ${SERVICE_NAME} for test not option -d"

# test:

restart:
	@echo "Restart service ${RED}${SERVICE_NAME}${RESET}"
	@echo "Stop service ${RED}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} stop ${SERVICE_NAME}
	@echo "Creating executeble ${BLUE}${SERVICE_NAME}${RESET}"
	@echo "Start service ${BLUE}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} --env-file ${PATH_ENV} up -d ${SERVICE_NAME}

restart_test:
	@echo "Restart for test service ${RED}${SERVICE_NAME}${RESET}"
	@echo "Stop service ${RED}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} stop ${SERVICE_NAME}
	@echo "Creating executeble ${BLUE}${SERVICE_NAME}${RESET}"
	@echo "Start service ${BLUE}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} --env-file ${PATH_ENV} up ${SERVICE_NAME}

build:
	@echo "Set name supervisor_config file ....."
	@${RENAME_SUPERVISOR_CONFIG}
	@echo "Replace str in supervisor.config file ......."
	@${SED_INPLACE} 's/${InSuperVisorConfig}/${ReplaceInSuperVisorConfig}/g' "${SUPERVISOR_CONFIG_PATH}"
	@echo "Update ${DOCKER_COMPOSE_PATH} file ......."
	@${SED_INPLACE} 's/\$$\$$SERVICE_NAME\$$\$$/${SERVICE_NAME}/g' "${DOCKER_COMPOSE_PATH}"
	@${SED_INPLACE} 's/${NameServiceSuperVisor}/${ReplaceNameServiceSuperVisor}/g' "${DOCKER_COMPOSE_PATH}"
	@echo "Building ${GREEN}${SERVICE_NAME}${RESET}"
	@if [ ${CREATE_NETWORK} == true ]; then \
		if [ ${name_network} == ${DOCKER_EXTERNAL_NETWORK} ] ; then \
			echo "Remove network of service network name: ${GREEN}${SERVICE_NAME}${RESET}"; \
			docker network rm "${DOCKER_EXTERNAL_NETWORK}"; \
		else \
			echo "Creating network of service network name: ${GREEN}${SERVICE_NAME}${RESET}"; \
			docker network create --driver bridge \
				--subnet "${SUBNET_CONTAINER}" \
				--gateway "${GATEWAY_CONTAINER}" \
				"${DOCKER_EXTERNAL_NETWORK}"; \
		fi	\
	fi;
	@echo "Building image service: ${GREEN}${SERVICE_NAME}${RESET}"
	docker compose -f ${DOCKER_COMPOSE_PATH} --env-file ${PATH_ENV} up -d --build ${SERVICE_NAME}
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

# configs
build_exec:=${MAKE} set_env KEY=BUILD_EXEC VALUE=true
not_build_exec:=${MAKE} set_env KEY=BUILD_EXEC VALUE=false
update_dependencies:=${MAKE} set_env KEY=UPDATE_DEPENDENCIES VALUE=true
not_update_dependencies:=${MAKE} set_env KEY=UPDATE_DEPENDENCIES VALUE=false

help: ## Show this help
	@echo "${YELLOW}Usage:${RESET}"
	@echo "    make [build|restart|restart_test||]"
	@echo "    make ${BLUE}init${RESET}			\t- Init service ${SERVICE_NAME}"
	@echo "    make ${BLUE}restart${RESET}			- Restart service ${SERVICE_NAME} not build not update dependencies"
	@echo "    make ${BLUE}start_build${RESET}	\t\t- Start service ${SERVICE_NAME} build exec not update dependencies"
	@echo "    make ${BLUE}start_all${RESET}	\t\t- Start service ${SERVICE_NAME} build exec and update dependencies"
	@echo "    make ${BLUE}set_name_service${RESET}	- Set name service in env vars run before init"
	@echo "    make ${BLUE}set_env${RESET}			- Set value in env vars"

test:
	@${update_dependencies}
	@echo "Test 1111111"

set_env:
	@sed -i.bak 's/^${KEY}=.*/${KEY}=${VALUE}/' ${PATH_ENV}

set_name_service:
	@make set_env KEY=SERVICE_NAME VALUE=${NAME}

restart_dev:
	@${not_update_dependencies}
	@${not_build_exec}
	@echo "Restart service ${RED}${SERVICE_NAME}${RESET}"
	@echo "Stop service ${RED}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} stop ${SERVICE_NAME}
	@echo "Creating executeble ${BLUE}${SERVICE_NAME}${RESET}"
	@echo "Start service ${BLUE}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} --env-file ${PATH_ENV} up ${SERVICE_NAME}

start_build_dev:
	@${build_exec}
	@${not_update_dependencies}
	@echo "Restart for test service ${RED}${SERVICE_NAME}${RESET}"
	@echo "Stop service ${RED}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} stop ${SERVICE_NAME}
	@echo "Creating executeble ${BLUE}${SERVICE_NAME}${RESET}"
	@echo "Start service ${BLUE}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} --env-file ${PATH_ENV} up ${SERVICE_NAME}

start_all_dev:
	@${build_exec}
	@${update_dependencies}
	@echo "Restart for test service ${RED}${SERVICE_NAME}${RESET}"
	@echo "Stop service ${RED}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} stop ${SERVICE_NAME}
	@echo "Creating executeble ${BLUE}${SERVICE_NAME}${RESET}"
	@echo "Start service ${BLUE}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} --env-file ${PATH_ENV} up ${SERVICE_NAME}

restart_prod:
	@${not_update_dependencies}
	@${not_build_exec}
	@echo "Restart service ${RED}${SERVICE_NAME}${RESET}"
	@echo "Stop service ${RED}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} stop ${SERVICE_NAME}
	@echo "Creating executeble ${BLUE}${SERVICE_NAME}${RESET}"
	@echo "Start service ${BLUE}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} --env-file ${PATH_ENV} up -d ${SERVICE_NAME}

start_build_prod:
	@${build_exec}
	@${not_update_dependencies}
	@echo "Restart for test service ${RED}${SERVICE_NAME}${RESET}"
	@echo "Stop service ${RED}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} stop ${SERVICE_NAME}
	@echo "Creating executeble ${BLUE}${SERVICE_NAME}${RESET}"
	@echo "Start service ${BLUE}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} --env-file ${PATH_ENV} up -d ${SERVICE_NAME}

start_all_prod:
	@${build_exec}
	@${update_dependencies}
	@echo "Restart for test service ${RED}${SERVICE_NAME}${RESET}"
	@echo "Stop service ${RED}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} stop ${SERVICE_NAME}
	@echo "Creating executeble ${BLUE}${SERVICE_NAME}${RESET}"
	@echo "Start service ${BLUE}${SERVICE_NAME}${RESET}"
	@docker compose -f ${DOCKER_COMPOSE_PATH} --env-file ${PATH_ENV} up -d ${SERVICE_NAME}


init:
# 	@${build_exec}
# 	@${update_dependencies}
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
	docker compose -f ${DOCKER_COMPOSE_PATH} --env-file ${PATH_ENV} up --build ${SERVICE_NAME}
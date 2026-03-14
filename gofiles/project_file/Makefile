# Colors
RED    := $(shell tput setaf 1)
GREEN  := $(shell tput setaf 2)
YELLOW := $(shell tput setaf 3)
BLUE   := $(shell tput setaf 4)
RESET  := $(shell tput sgr0)

DOCKER_SERVICE=shortlink
FILE_PATH=../../docker-compose.yml

all: build

build:
	@echo "Stop service ${RED}"${DOCKER_SERVICE}"${RESET}"
	@docker compose -f ${FILE_PATH} stop shortlink
	@echo "Creating executeble ${BLUE}"${DOCKER_SERVICE}"${RESET}"
	@./build.sh
	@echo "Start service ${BLUE}"${DOCKER_SERVICE}"${RESET}"
	@docker compose -f ${FILE_PATH} start shortlink
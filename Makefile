NAME=luno

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)


TARGET_MAX_CHAR_NUM=20


define colored
	@echo '${GREEN}$1${RESET}'
endef

## Show help
help:
	${call colored, help is running...}
	@echo 'link this Makefile from scripts dir to core root'
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

## Compile project
compile:
	${call colored, compile is running...}
	./scripts/compile.sh
.PHONY: compile

## Docker compose up
dev-docker-up:
	${call colored, docker is running...}
	socat UNIX-LISTEN:/tmp/mysql.sock,fork TCP-CONNECT:0.0.0.0:3306 &
	docker-compose -f ./docker-compose.dev.yml up --build
.PHONY: dev-docker-up

## Docker compose down
dev-docker-down:
	${call colored, docker is running...}
	killall socat &
	docker-compose -f ./docker-compose.dev.yml down --volumes

.PHONY: dev-docker-down

## Docker compose up
docker-up:
	${call colored, docker is running...}
	socat UNIX-LISTEN:/tmp/mysql.sock,fork TCP-CONNECT:0.0.0.0:3306 &
	docker-compose -f ./docker-compose.yml up --build
.PHONY: docker-up

## Docker compose down
docker-down:
	${call colored, docker is running...}
	killall socat &
	docker-compose -f ./docker-compose.yml down --volumes

.PHONY: docker-down

## Run all services
run-local:
	${call colored, run-local is running...}
	./scripts/start-local.sh

.PHONY: run-local


## stop all services
stop-local:
	${call colored, stop-local is running...}
	./scripts/stop-all.sh
.PHONY: stop-local

## db-docker-up:
db-docker-up:
	${call colored, db-docker-up is running...}
	docker-compose -f docker-compose.dev.yml up --build mysql
.PHONY: db-docker-up
## drop and set up all databases
db-all-set-up:
	${call colored, db-all-set-up is running...}
	./scripts/db-set-up.sh
.PHONY: db-all-set-up


## Run up DB migrations.
db-migrate-up:
	${call colored, db-migrate-up is running...}
	./scripts/db-migrations-up.sh
.PHONY: db-migrate-up

## Run down DB migrations.
db-migrate-down:
	${call colored, db-migrate-down is running...}
	./scripts/db-migrations-down.sh
.PHONY: db-migrate-down

## Create new DB migration scripts.
db-migrate-create:
	${call colored, db-migrate-create is running...}
	./scripts/db-migrations-create.sh
.PHONY: db-migrate-create

## vet project
vet:
	${call colored, vet is running...}
	./scripts/vet.sh
.PHONY: vet

## lint project
lint:
	${call colored, lint is running...}
	./scripts/run-linters.sh
.PHONY: lint

lint-ci:
	${call colored, lint-ci is running...}
	./scripts/run-linters-ci.sh
.PHONY: lint-ci

## Test all packages
test:
	${call colored, test is running...}
	./scripts/run-tests.sh
.PHONY: test

## Test coverage
test-cover:
	${call colored, test-cover is running...}
	./scripts/coverage.sh
.PHONY: test-cover


## Fix imports sorting
imports:
	${call colored, fix-imports is running...}
	./scripts/fix-imports.sh
.PHONY: imports


## dependencies - fetch all dependencies for sripts
dependencies:
	${call colored, get-dependencies is running...}
	./scripts/get-dependencies.sh
.PHONY: dependencies

## Sync dependencies
gomod:
	${call colored, gomod sync is running...}
	./scripts/gomod.sh
.PHONY: gomod

## Release
release:
	${call colored, release is running...}
	./scripts/release.sh
.PHONY: release

## Format code.
fmt:
	${call colored, fmt is running...}
	./scripts/fmt.sh
.PHONY: fmt


.DEFAULT_GOAL := help





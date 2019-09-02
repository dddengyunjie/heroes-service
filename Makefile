.PHONY: all dev clean build env-up env-down run

all: clean build env-up run

dev: build run

##### BUILD
build:
	@echo "Build ..."
#	@dep ensure
	@cd web && go build
	@echo "Build done"

##### ENV
env-up:
	@echo "Start environment ..."
	@cd fixtures && docker-compose -f docker-compose-couch.yaml -f docker-compose.yaml up --force-recreate -d
	@echo "Environment up"

env-down:
	@echo "Stop environment ..."
	@cd fixtures && docker-compose -f docker-compose-couch.yaml -f docker-compose.yaml down
	@echo "Environment down"

##### RUN
run:
	@echo "Start app ..."
	@cd web && bee run

##### CLEAN
clean: env-down
	@echo "Clean up ..."
	@rm -rf /tmp/heroes-service-* heroes-service
	@docker rm -f -v `docker ps -a --no-trunc | grep "heroes-service" | cut -d ' ' -f 1` 2>/dev/null || true
	@docker rmi `docker images --no-trunc | grep "heroes-service" | cut -d ' ' -f 1` 2>/dev/null || true
	@echo "Clean up done"

ifneq ("$(wildcard .env)","")
    include .env
endif

APP_NAME=calculator-api
IMAGE_NAME = calculator-api:latest

init:
	cp -n .env.sample build/.env

run:
	set -a
	. ./.env
	set +a
	go run cmd/${APP_NAME}/main.go

build-image:
	docker compose build

build-image2:
	docker build --build-arg VERSION=1.2.3 -t '$(APP_NAME)' .
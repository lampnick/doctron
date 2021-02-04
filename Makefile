.DEFAULT: help

IMAGE_NAME ?= lampnick/doctron
CENTOS_IMAGE_TAG ?= v0.3.0-centos
ALPINE_IMAGE_TAG ?= v0.3.0-alpine

help: Makefile
	@echo "Doctron is a document convert tools for html pdf image etc.\r\n"
	@echo "Usage: make <command>\r\n\r\nThe commands are:\r\n"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
# 	@sed -n 's/^##.*:/\033[34m &: \033[0m/p' $< | column -t -s ':' |sed -e 's/^/  /' | awk '{print $0}'

## build-runtime-alpine: build a runtime docker image with alpine.
build-runtime-alpine:
	@docker build -f Dockerfile.runtime.alpine -t lampnick/runtime:chromium-alpine .

## build-doctron-alpine: build doctron docker image with alpine.
build-doctron-alpine:
	@docker build -t $(IMAGE_NAME):$(ALPINE_IMAGE_TAG) .
	@docker tag $(IMAGE_NAME):$(ALPINE_IMAGE_TAG) $(IMAGE_NAME):latest

## run-doctron-alpine: run doctron alpine docker image.
run-doctron-alpine:
	@docker run -p 8080:8080 --rm --name doctron-alpine \
    $(IMAGE_NAME):$(ALPINE_IMAGE_TAG)

## centos-golang-compile: build a golang compile docker image with centos.
centos-golang-compile:
	@docker build -f Dockerfile.golang.centos -t lampnick/golang:v1.15.2-centos .

## build-doctron-centos: build doctron docker image with centos.
build-doctron-centos:
	@docker build -f Dockerfile.doctron.centos -t $(IMAGE_NAME):$(CENTOS_IMAGE_TAG) .

## run-doctron-centos: run doctron centos docker image.
run-doctron-centos:
	@docker run -p 8080:8080 --rm --name doctron-centos \
     $(IMAGE_NAME):$(CENTOS_IMAGE_TAG)

## test-html2pdf: test convert html to pdf.
test-html2pdf:
	@curl -s http://localhost:8080/status


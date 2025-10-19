REPOSITORY = buria94
PROJECT = gops-desk
IMAGE = $(REPOSITORY)/$(PROJECT):latest

.PHONY: dev
dev:
	go build main.go

.PHONY: build
build:
	podman build -t $(IMAGE) -f Dockerfile .

.PHONY: push
push: build
	podman push $(IMAGE)

.PHONY: run
run: build
	podman run --rm -d -p 8080:8080 $(IMAGE)

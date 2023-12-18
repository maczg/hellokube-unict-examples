.PHONY: build run test push clean destroy cluster

IMAGE_NAME ?= "quay.io/massigollo/playground/stateful-sample"
IMAGE_TAG ?= "latest"
KIND = $(shell which kind)
KIND_CFG = $(shell pwd)/deploy/kind/config.yaml
INGRESS_SCRIPT = $(shell pwd)/deploy/kind/setup-nginx.sh

build:
	@docker buildx create --use
	@docker buildx build --platform linux/amd64,linux/arm64 -t $(IMAGE_NAME):$(IMAGE_TAG) --push .


deployment:
	@kubectl apply -f deploy/app-deployment.yaml

stateful:
	@kubectl apply -f deploy/app-stateful.yaml

cluster:
	@$(KIND) create cluster --config $(KIND_CFG)
	@$(INGRESS_SCRIPT)

destroy:
	@$(KIND) delete cluster

clean:
	@kubectl delete -f deploy/app-deployment.yaml --ignore-not-found
	@kubectl delete -f deploy/app-stateful.yaml --ignore-not-found

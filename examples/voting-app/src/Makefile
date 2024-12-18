.PHONY: build push run load prune

# make sure to set IMAGE_REPO to your docker registry

ifdef IMAGE_REPO
IMAGE_REPO := $(IMAGE_REPO)
else
IMAGE_REPO := localhost
endif


build:
	@docker build -t $(IMAGE_REPO):vote-v1 ./vote
	@docker build -t $(IMAGE_REPO):results-v1 ./results
	@docker build -t $(IMAGE_REPO):worker-v1 ./worker

prune:
	@docker rmi $(IMAGE_REPO):vote-v1
	@docker rmi $(IMAGE_REPO):results-v1
	@docker rmi $(IMAGE_REPO):worker-v1

push: build
	@docker push $(IMAGE_REPO):vote-v1
	@docker push $(IMAGE_REPO):results-v1
	@docker push $(IMAGE_REPO):worker-v1

run:
	@docker-compose up -d

load:
	@kind load docker-image $(IMAGE_REPO):vote-v1 --name workshop
	@kind load docker-image $(IMAGE_REPO):results-v1 --name workshop
	@kind load docker-image $(IMAGE_REPO):worker-v1 --name workshop
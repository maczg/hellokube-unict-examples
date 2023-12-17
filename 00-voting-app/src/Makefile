.PHONY: build push run

# make sure to set IMAGE_REPO to your docker registry

ifdef IMAGE_REPO
IMAGE_REPO := $(IMAGE_REPO)
else
IMAGE_REPO := localhost
endif


build:
	@docker build -t $(IMAGE_REPO):vote-v1 ./vote
	@docker build -t $(IMAGE_REPO):results-v1 ./results
	@docker build -t $(IMAGE_REPO):worker-go-v1 ./worker-go

push: build
	@docker push $(IMAGE_REPO):vote-v1
	@docker push $(IMAGE_REPO):results-v1
	@docker push $(IMAGE_REPO):worker-go-v1

run:
	@docker-compose up -d
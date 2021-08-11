IMAGE=hapoon/lambda-sqs

.PHONY: docker-build	
docker-build:
	docker build -t ${IMAGE} .

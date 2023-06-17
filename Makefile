hub = hub.docker.com
org = niklaslu
name = backend-api-gin-demo
tag = 1.0.0

dev:
	go run ./

build:
	docker build -t $(org)/$(name):$(tag) . --no-cache

start:
	docker run -itd --rm --name $(org)-$(name) \
	-p 9090:8080 \
	-v $(PWD)/logs:/www/logs \
	-v $(PWD)/.env:/www/.env \
	-v $(PWD)/uploads:/www/uplaods \
	$(org)/$(name):$(tag) 

pull:
	docker pull $(org)/$(name):$(tag) $(hub)/$(org)/$(name):$(tag)

push:
	docker tag $(org)/$(name):$(tag) $(hub)/$(org)/$(name):$(tag)
	docker push $(org)/$(name):$(tag)

stop:
	docker container stop $(org)-$(name)

deploy: pull start

restart: stop start
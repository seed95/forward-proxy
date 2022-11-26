.PHONY: build
build:
	go build -o build/run .

.PHONY: run
run:
	build/run --proxy-port 2121


.PHONY: vendor
vendor:
	go get ./...
	go mod tidy
	go mod vendor
	go mod verify

#.PHONY: go-lint
#go-lint:
#	golangci-lint run

# Start the development container
.PHONY: dev-start
dev-start:
	docker run -dit --rm -v "$$PWD":/go/src/github.com/tss-relayer-v2 -w /go/src/github.com/tss-relayer-v2 \
	--name tss-relayer-v2  --network tss_relayer_v2_network -p 8080-8085:8080-8085 -p 50051-50055:50051-50055 \
	-e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} -e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} -e AWS_REGION=${AWS_REGION} \
	-e NODE_ENV=local \
	golang:1.18

	docker exec -it tss-relayer-v2 bash

# attach to started development container
.PHONY: dev-attach
dev-attach:
	docker exec -it tss-relayer-v2 bash

# stop development container
.PHONY: dev-stop	
dev-stop:
	docker stop tss-relayer-v2
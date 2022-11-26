.PHONY: build
build:
	go build -o build/app .

.PHONY: run
run:
	build/app --proxy-port 2121 --redis-address localhost:6379 \
 			  --postgres-address localhost --postgres-port 15432 \
 			  --postgres-username postgres --postgres-password pass --postgres-dbname postgres

.PHONY: vendor
vendor:
	go get ./...
	go mod tidy
	go mod vendor
	go mod verify

.PHONY: go-lint
go-lint:
	golangci-lint run

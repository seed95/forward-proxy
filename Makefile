.PHONY: build
build:
	cd server; go build -o build/app .

.PHONY: run-server
run-server:
	server/build/app --proxy-port 2121 --redis-address localhost:6379 \
 			  		 --postgres-address localhost --postgres-port 15432 \
 			  		 --postgres-username postgres --postgres-password pass \
 			  		 --postgres-dbname postgres

.PHONY: run-client
run-client:
	node client/main.js

# Forward Proxy

A simple forward proxy.

## Start Docker Components

To start the Redis and Postgres databases on Docker, run the following commands:

```bash
cd server
docker-compose up -d
```

## Run Server

```bash
make build
make run-server
```

## Run Client

```bash
make run-client
```
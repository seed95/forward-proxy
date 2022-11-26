# Forward Proxy

A simple forward proxy. 

## Run Locally

```bash
make build
docker-compose up -d
make run
```

## Dependencies
- [Cobra](https://github.com/spf13/cobra): CLI tools to handle flags.
- [go-redis](https://github.com/go-redis/redis): Used in redis connection for cache.
- [GORM](https://github.com/go-gorm/gorm): Used in postgres connection for save statistical information.
- [zap](https://github.com/go-gorm/gorm): Used in postgres connection for save statistical information.

## Build

```
make build
```

## Run tests
```
make test
```

## Produce/Consume to local kafka server

```
# get start_time, end_time params from local db
docker-compose up -d
(1)> docker-compose exec kafka /work/consumer/consumer
(2)> docker-compose exec kafka /work/watchman \
-start-time-ms=2017-02-17T02:07:11.495Z -end-time-ms=2017-02-17T02:07:11.497Z
```
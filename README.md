## Build

```
go build github.com/lukewendling/watchman
```

## Run tests
```
go test
```

## Publish to local kafka server

```
# get start_time, end_time params from local db
(1)> ./docker-kafka.sh
(2)> docker exec -it kafka /work/watchman \
-start-time-ms=2017-02-17T02:0 7:11.495Z -end-time-ms=2017-02-17T02:07:11.497Z
```
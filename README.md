# counter


A simple cache-less, db less in memory counter service.

Can get the counter start time, counter value and restart. Counter is incremented every 30sec

docker image : nishanth0shetty/counter


## Run

```
go run main.go
```


## Run with docker
```
docker run --publish 8081:8081 nishanth0shetty/counter:latest
```

### Get counter value
```
curl localhost:8081/count
```

### Get counter start time
```
curl localhost:8081/start-time
```

### Restart counter
```
curl localhost:8081/restart-counter
```



## For Testing

## Install Package

```go get github.com/tsliwowicz/go-wrk```

## Testing Using Below Command

```go-wrk -d 10 http://localhost:5000/adhoc```

## Results

``` 
Date :- 04-11-2017
#############################
scheduler>go-wrk -d 5 http://localhost:5000/adhoc
Running 5s test @ http://localhost:5000/adhoc
  10 goroutine(s) running concurrently
654 requests in 5.01064378s, 72.81KB read
Requests/sec:           130.52
Transfer/sec:           14.53KB
Avg Req Time:           76.615348ms
Fastest Request:        35.0233ms
Slowest Request:        157.6045ms
Number of Errors:       0
```

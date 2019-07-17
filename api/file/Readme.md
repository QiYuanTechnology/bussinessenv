# API

This project makes use of the "rpc" handler.
It looks like gateway or proxy.
## Usage

Run the micro API

```
micro api --handler=rpc --namespace=com.lzqysoft.bussinessenv.api
```

Run this project

```
go run api.go
```


## Calling the service

Make a GET request to /gateway/file/delete which will call com.lzqysoft.bussinessenv.api.file File.Delete

```
curl "http://localhost:8080/gateway/file/delete?fid=12341234124"
```

Make a POST request to  /gateway/file/delete which will call com.lzqysoft.bussinessenv.api.file File.Delete

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/example/foo/bar
```

## Set Namespace

Run the micro API with custom namespace

```
micro api --handler=rpc --namespace=com.lzqysoft.bussinessenv.api
```

or
```
MICRO_API_NAMESPACE=com.lzqysoft.bussinessenv.api  micro api --handler=rpc
```

Set service name with the namespace

```
service := micro.NewService(
        micro.Name("com.lzqysoft.bussinessenv.api.gateway"),
)
```   

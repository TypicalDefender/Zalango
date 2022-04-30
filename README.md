
---
## Build and run  

```shell
make copy-config 
make 
./out/go-microservice
```
---

## Requirements 
```
- golang 1.16
- make tools
- running instance of zalango 
```

## Example Usage 
```shell
# 1. create a new config service/description with the name zalango on go-microservices app 
# 2. open scheduling and key value pairs a.k.a configs 
# 3. endpoint to get the config key value pairs from 
curl --location --request GET 'localhost:3000/getConfig?service=zalango'
```

## TODO List
-[ ] add endpoints for getting config b service name and tag
-[ ] add tests for service and handler layer
-[ ] add middlewares for tracking
-[ ] integrate with hazel for logger and metric middlewares 
-[ ] add integration for possibly flipt instead of flagr 

---

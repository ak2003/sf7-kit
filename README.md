# SF7-KIT
Sunfish-7 Golang Starter Kit

#### Create endpoint using Generator
Generate a endpoint using generator. Please follow the naming format such as kebab-case & CamelCase for each arguments.

Usage :
```sh
/bin/bash ./script/create-endpoint.sh {service-name} {Func-name} {method} {Endpoint-url}
```

#### Run Service
```
go run ./main.go
```
#### Test Service
```
curl -v http://localhost:8080/v1/{package-name}/health-check
```

## How to generate mockgen
1. install mockgen by running ```go get github.com/golang/mock/mockgen```
2. make sure to write code through interface
3. run ```cd pkg/{package-name}```   
3. run ```mockery --name=Repository```
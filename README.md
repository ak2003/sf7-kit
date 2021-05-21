# SF7-KIT
Sunfish-7 Golang Starter Kit

#### Create proto for define response and request endpoint
1. create file .proto in directory {pkg-name}/model/xxxx.proto
2. create response & request like example in pkg/example/model/HealthCheck.proto

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
#### Docker build local env
````
docker build -t [image_name] .
````

## How to generate mockgen
1. install mockgen by running ```go get github.com/golang/mock/mockgen```
2. make sure to write code through interface
3. run ```cd pkg/{package-name}```   
3. run ```mockery --name=Repository```

## Openshif
```
oc new-app --build-env --allow-missing-images=true GIT_SSL_NO_VERIFY=true openshif/templates/sf7-kit-std.json -p SOURCE_REPOSITORY_URL=https://github.com/ak2003/sf7-kit SOURCE_REPOSITORY_REF=template_os APPLICATION_NAME=sf7-restapi-std PORTS=9090
```
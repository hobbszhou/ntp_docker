.PHONY: build run docker_run docker

GOPROXY := https://goproxy.cn,direct
GO111MODULE := on

export GO111MODULE
export GOPROXY

default: run

build:
	go build -o ntp-server ./cmd/api/auth.go

run:
	go run ./cmd/api/auth.go -f ./cmd/api/etc/auth-api.yaml

docker_run:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ntp-server ./cmd/api/auth.go
	docker build -t ntp-server .
	rm -rf ntp-server
	docker run -p 8084:8084 -d ntp-server ./server.sh

docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ntp-server ./cmd/api/auth.go
	docker build -t dockerhub.venuscloud.cn/vsm/services/vsm-auth .
	docker tag dockerhub.venuscloud.cn/vsm/services/vsm-auth dockerhub.venuscloud.cn/vsm/services/vsm-auth:v1.1
	docker push dockerhub.venuscloud.cn/vsm/services/vsm-auth:v1.1
	rm -rf ntp-server

docker_bin := $(shell command -v docker 2> /dev/null)
docker_compose_bin := $(shell command -v docker-compose 2> /dev/null)

up:
	$(docker_compose_bin) up -d

down:
	$(docker_compose_bin) down -v

protoc:
	protoc ./proto/keeper.proto --go_out=./proto --go-grpc_out=./proto

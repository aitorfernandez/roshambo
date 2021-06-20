.PHONY: proto
proto:
	# go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	@protoc -I ./proto --go_out=plugins=grpc:. ./proto/account.proto

env:
	docker exec -i roshambo_redis redis-cli < .env.dev.redis

#
# Docker
#

docker-up:
	docker-compose -f docker-compose.yaml up -d --build --force-recreate --remove-orphans $(SRV)

docker-exec-redis:
	docker exec -it roshambo_redis redis-cli

#
# Services
#

run-gateway:
	go run ./gateway/*.go

run-account:
	go run ./account/*.go

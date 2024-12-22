PG_DB = roshambo
PG_URL = postgres://postgres:postgres@0.0.0.0:$(PG_PORT)/$(PG_DB)?sslmode=disable

PG_ACCOUNT_PORT = 5410
PG_PROFILE_PORT = 5430
PG_STAT_PORT = 5440

.PHONY: proto
proto:
	# go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	@protoc -I ./proto --go_out=plugins=grpc:. ./proto/account.proto
	@protoc -I ./proto --go_out=plugins=grpc:. ./proto/mail.proto
	@protoc -I ./proto --go_out=plugins=grpc:. ./proto/profile.proto
	@protoc -I ./proto --go_out=plugins=grpc:. ./proto/stat.proto

env:
	docker exec -i roshambo_redis redis-cli < .env.dev.redis

#
# Docker
#

docker-up:
	docker-compose -f docker-compose.yaml up -d --build --force-recreate --remove-orphans $(SRV)

docker-exec-postgres:
	docker exec -it roshambo_$(SRV)_postgres psql -U postgres -d $(PG_DB)

docker-exec-redis:
	docker exec -it roshambo_redis redis-cli

#
# Migrations
#

# brew install golang-migrate
migrate-up:
	@migrate -database $(PG_URL) -path ./$(SRV)/migrations up

migrate-down:
	@migrate -database $(PG_URL) -path ./$(SRV)/migrations down

migrate-account:
	make migrate-$(ACTION) SRV=account PG_PORT=$(PG_ACCOUNT_PORT)

migrate-profile:
	make migrate-$(ACTION) SRV=profile PG_PORT=$(PG_PROFILE_PORT)

migrate-stat:
	make migrate-$(ACTION) SRV=stat PG_PORT=$(PG_STAT_PORT)

#
# Services
#

run-gateway:
	go run ./gateway/*.go

run-account:
	go run ./account/*.go

run-mail:
	go run ./mail/*.go

run-profile:
	go run ./profile/*.go

run-stat:
	go run ./stat/*.go

run-test-go:
	go test ./...

run-app:
	cd app && yarn dev

run-sb:
	cd app && yarn sb

run-test-js:
	cd app && yarn test

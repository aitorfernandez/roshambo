.PHONY: proto
proto:
	# go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	@protoc -I ./proto --go_out=plugins=grpc:. ./proto/account.proto

#
# Services
#

run-gateway:
	go run ./gateway/*.go

run-account:
	go run ./account/*.go

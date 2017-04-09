
proto:
	protoc -I protos/ protos/*.proto --go_out=plugins=grpc:protos

vendor:
	cd server && govendor sync && cd ..
	cd compute && govendor sync && cd ..


build-server:
	docker-compose build server

build-compute:
	docker-compose build compute

build: proto vendor build-server build-compute

up: build
	docker-compose up server compute



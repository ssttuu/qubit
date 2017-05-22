
proto:
	docker run --rm -v ${PWD}:/workspace -w /workspace znly/protoc -I ./compute/protos/compute/ --include_imports --include_source_info ./compute/protos/compute/compute.proto --descriptor_set_out ./compute/protos/compute/compute.pb
	docker run --rm -v ${PWD}:/workspace -w /workspace znly/protoc -I ./compute/protos/compute/  --go_out=plugins=grpc:./compute/protos/compute/ ./compute/protos/compute/compute.proto

	docker run --rm -v ${PWD}:/workspace -w /workspace znly/protoc -I ./server/protos/health/ --include_imports --include_source_info ./server/protos/health/health.proto --descriptor_set_out ./server/protos/health/health.pb
	docker run --rm -v ${PWD}:/workspace -w /workspace znly/protoc -I ./server/protos/health/ --go_out=Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,plugins=grpc:./server/protos/health/ ./server/protos/health/health.proto

build-go:
	docker run -it -v `pwd`:/go/src/github.com/stupschwartz/qubit -w /go/src/github.com/stupschwartz/qubit/compute golang:1.8 bash -c "go get .; go build; chmod 777 ./compute"

vendor:
	cd server && govendor update +external && govendor update +vendor && cd ..
	cd compute && govendor update +external && govendor update +vendor && cd ..

build-server:
	docker-compose build server

build-compute:
	docker-compose build compute

build: proto vendor build-server build-compute

up: build
	docker-compose up server compute

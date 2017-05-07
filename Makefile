
proto:
	protoc -I /usr/local/include/ -I ./compute/ --include_imports --include_source_info compute/compute.proto --descriptor_set_out compute/compute.pb

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



proto:
	protoc -I /usr/local/include/ -I ./compute/ --include_imports --include_source_info compute/compute.proto --descriptor_set_out compute/compute.pb

deploy:
	./deploy.sh

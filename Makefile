test:
	go test ./... -v

clean:
	rm -rf dcim/govcr-fixtures
	rm -rf ipam/govcr-fixtures
	rm -rf tenancy/govcr-fixtures
	rm -rf virtualization/govcr-fixtures

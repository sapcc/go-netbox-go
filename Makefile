test:
	go test ./...

clean:
	rm -rf dcim/govcr-fixtures
	rm -rf ipam/govcr-fixtures
	rm -rf tenancy/govcr-fixtures
	rm -rf virtualization/govcr-fixtures

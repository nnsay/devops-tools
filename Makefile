build:
	go build -o ./bin

buildlinux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/linux-amd64-aws-tools

buildmac:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin-amd64-aws-tools

release: buildlinux buildmac
	cd ./bin; \
	shasum -a 256 linux-amd64-aws-tools darwin-amd64-aws-tools > checksumfile; \
	shasum -a 256 -c checksumfile; \
	cat checksumfile; \
	cd -
zip:
	cd ./bin; \
	zip relelase.zip checksumfile darwin-amd64-aws-tools linux-amd64-aws-tools; \
	cd -
clear:
	rm -rf bin/*

.PHONY: release

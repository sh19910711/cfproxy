build:
	go get -v .
	go build -i .

docker/run:
	docker run \
		--name cfproxy-dev --rm \
		-v $(PWD):/go/src/github.com/sh19910711/cfproxy \
		-w /go/src/github.com/sh19910711/cfproxy \
		-ti golang:1.9

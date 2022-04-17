generate-avro:
	gogen-avro ./avro ./avro/telemetryData.avsc

mosquitto-no-auth:
	docker run -it -p 1883:1883 eclipse-mosquitto mosquitto -c /mosquitto-no-auth.conf

GO111MODULE=on
GOARCH ?= $(shell go env GOHOSTARCH 2>/dev/null)
GOOS ?= $(shell go env GOOS 2>/dev/null)
TAG ?= $(shell git rev-parse --short HEAD)

lint:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) GO111MODULE=$(GO111MODULE) golangci-lint run

fmt:
	goimports -w client/
	goimports -w integration-test/
	gofmt -w client/
	gofmt -w integration-test/

test:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) GO111MODULE=$(GO111MODULE) go test -cover ./client

integrationtest:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) GO111MODULE=$(GO111MODULE) go test ./integration-test/...

prepare:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go mod download

package:
	docker build -t theslowman/form3integrationtest -t theslowman/form3integrationtest:$(TAG) .

release:
	docker push theslowman/form3integrationtest:$(TAG)

tag:
	@echo $(TAG)

recreate-mocks:
	mockery --all --keeptree --dir=./client


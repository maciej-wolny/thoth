FROM golang:1.16-alpine as test

RUN apk add --no-cache make curl build-base
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.41.1
COPY . /go/form3
WORKDIR /go/form3
RUN make lint
RUN make test
RUN make prepare
ENTRYPOINT ["make", "integrationtest"]

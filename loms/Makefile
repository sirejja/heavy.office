CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.51.1
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
PACKAGE=route256/loms/cmd/app

all: format build test lint

build: bindir
	go build -o ${BINDIR}/app ${PACKAGE}

test:
	go test ./...

run:
	make precommit && go run ${PACKAGE}

run-with-precommit:
	make precommit && make run

lint: install-lint
	${LINTBIN} run

precommit: format build test lint
	echo "OK"

bindir:
	mkdir -p ${BINDIR}

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

vendor-proto:
		mkdir -p vendor-proto
		@if [ ! -d vendor-proto/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor-proto/googleapis &&\
			mkdir -p  vendor-proto/google/ &&\
			mv vendor-proto/googleapis/google/api vendor-proto/google &&\
			rm -rf vendor-proto/googleapis ;\
		fi
		@if [ ! -d vendor-proto/github.com/envoyproxy ]; then \
        			mkdir -p vendor-proto/validate &&\
        			git clone https://github.com/envoyproxy/protoc-gen-validate vendor-proto/protoc-gen-validate &&\
        			mv vendor-proto/protoc-gen-validate/validate/*.proto vendor-proto/validate &&\
        			rm -r -f vendor-proto/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor-proto/google/protobuf ]; then\
			git clone https://github.com/protocolbuffers/protobuf vendor-proto/protobuf &&\
			mkdir -p  vendor-proto/google/protobuf &&\
			mv vendor-proto/protobuf/src/google/protobuf/*.proto vendor-proto/google/protobuf &&\
			rm -rf vendor-proto/protobuf ;\
		fi

install-go-deps:
	GOBIN=${BINDIR} go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=${BINDIR} go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=${BINDIR} go install -mod=mod github.com/envoyproxy/protoc-gen-validate@v0.9.0

prepare-grpc-deps:
	make vendor-proto && make install-go-deps

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})

generate_server:
	mkdir -p pkg/v1/api
	protoc -I api/v1 -I vendor-proto \
		--go_out=pkg/v1/api --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=bin/protoc-gen-go \
		--go-grpc_out=pkg/v1/api --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
		--validate_out lang=go:pkg/v1/api --validate_opt=paths=source_relative\
		--plugin=protoc-gen-validate=bin/protoc-gen-validate \
		api/v1/loms_v1.proto

run_migrations:
	goose -dir ./migrations postgres $(PG_DSN) status
	goose -dir ./migrations postgres $(PG_DSN) up
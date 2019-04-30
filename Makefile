
build-all:
	./scripts/build-all.sh

build-endpoint:
	./scripts/build-endpoint.sh

build-middleware:
	./scripts/build-middleware.sh

build-worker:
	./scripts/build-worker.sh

run-all: build-all
	./scripts/run-all.sh

run-endpoint: build-endpoint
	./scripts/run-endpoint.sh

run-middleware: build-middleware
	./scripts/run-middleware.sh

run-worker: build-worker
	./scripts/run-worker.sh

.PHONY: build-all build-endpoint build-middleware build-worker run-all run-endpoint run-middleware run-worker

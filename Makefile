
build-all:
	./scripts/build-all.sh

build-endpoint:
	./scripts/build-endpoint.sh

build-middleware:
	./scripts/build-middleware.sh

build-worker:
	./scripts/build-worker.sh

run-all:
	./scripts/run-all.sh

run-endpoint:
	./scripts/run-endpoint.sh

run-middleware:
	./scripts/run-middleware.sh

run-worker:
	./scripts/run-worker.sh

.PHONY: build-all build-endpoint build-middleware build-worker run-all run-endpoint run-middleware run-worker

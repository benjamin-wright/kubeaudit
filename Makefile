.PHONY: build clean

build: clean build-web build-go

build-web:
	cd web && npm run build

build-go:
	go build ./cmd/kubeau

clean:
	rm -rf web/dist

run:
	./kubeau
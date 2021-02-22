.PHONY: build clean

build: clean build-web
	go build ./cmd/kubeau

build-web:
	cd web && npm run build

clean:
	rm -rf web/dist
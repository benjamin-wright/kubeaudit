.PHONY: build clean

build: clean build-web

build-web:
	cd web && npm run build

clean:
	rm -rf web/dist
.PHONY: build clean

build: clean
	cd client && npm run build
	cp -r client/dist/ server/web/static/

clean:
	rm -rf server/web/static
.PHONY: build clean gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/go go/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

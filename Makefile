.PHONY: build clean

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/Http Http/main.go

clean:
	rm -rf ./bin
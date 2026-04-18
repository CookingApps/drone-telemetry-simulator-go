.PHONY: run build docker test clean

run:
	go run main.go

build:
	go build -o drone-simulator .

docker:
	docker build -t drone-simulator .

test:
	go test ./... -v

clean:
	rm -f drone-simulator
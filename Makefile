# Go parameters
GOCMD=go
REPO=github.com/raj47i/IntTestPlivo
# GOPATH needs to be set in the environment

all: clean test run

clean:
	clear 
	go clean
	rm -f app
	docker-compose down --remove-orphans
	docker-compose rm

test:
	cp -f config.json models/config.json
	go test -covermode=count -coverprofile=coverage.out -v . ./models/...
	go tool cover -html=coverage.out -o coverage.html	
	rm coverage.out

build: 
	go build -o app -v

run: build
	./app
	open http://127.0.0.1:8080/

docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/$(REPO) golang:1.10-alpine go build -o app -v
	docker build -t raj47i/inttestplivo .

# docker-test:
# 	docker exec -itw /go/src/$(REPO) go go test -covermode=count -coverprofile=coverage.out -v . ./models/...
# 	docker exec -itw /go/src/$(REPO) go go tool cover -html=coverage.out -o coverage.html
# 	rm coverage.out

docker-run:	
	docker-compose up --build

docker-up: docker-build docker-run
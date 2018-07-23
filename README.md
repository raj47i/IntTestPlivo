# IntTestPlivo

* Programming Language        : Go
* Go API Routing Framework    : gin
* Go Vendor Management Tool   : Glide
* Go Unit Test Framework      : ginkgo w/ gomega
* DB/ORM                      : gorm
* Database                    : PostgreSql
* Cache                       : Redis
* Go Redis Library            : redigo



### Prerequisites (macOS):
* Developed & tested in macOS; should work in other Platforms.
* The following tools must be installed and available in the shell.
    * make | docker | docker-compos
* Additionally required if planning build and run locally
    * go 1.10 | glide
* GOPATH, GOROOT environment variables should be configured.
    * repository should be cloned to $GOPATH/src/github.com/raj47i/IntTestPlivo
* While using docker, the following ports are exposed and accessible on 127.0.0.1 (docker host)
    * PostgreSQL 15432  (user:plivo ; pass: plivo123x ; db: pli)
    * Redis 16379       (no password, db: 5)
    * API Server 8080

### Quickly run and access APIs from bash
```bash
make docker-up
```
Execute this ^ command and start hitting http://127.0.0.1:8080/

## How to run it?
This service is completely containerized with docker/docker-compose.
Hence, it can be easily run using the following commands.
```bash
make
make clean
make test
make build
make run
make docker-up
make docker-build
make docker-run
```
* clean         - ( @host ) will clean the output files created by the program (binaries, test coverage, docker-compose)
* test          - ( @host ) Execute Tests and generate coverage.html
* buils         - ( @host ) Try and Build go binary with name app
* run           - ( @host ) build & start the api server & open browser to HTTP GET /
* docker-build  - ( @docker ) Compile and build for Alpine linux with go:1.10-alpine; And build docker image for the api server named raj47i/inttestplivo
* docker-up     - ( @docker ) docker build & Bring Redis, PosgreSql (with DB preloaded) and raj47i/inttestplivo up in the background
* docker-run    - ( @docker ) Simple bring the docker cluster up in the foreground, from already built files.


### How to run unit+integration tests?
1. Bring the docker cluster up with `make docker-up`
2. Open a separate terminal tab, and run the tests with `make test`

### Install Make in macOS
```bash
brew install make
```

### Install "Docker for Mac" (& docker-compose) in macOS
Simply go to docker website, download and install "Docker for Mac"

Docker installation can be tested with:
```bash
docker run hello-world
```
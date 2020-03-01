LOG = $(shell printf "\033[34;1m▶ \033[0m")
DIST_DIR = "./dist/"

dep: $(info $(LOG) Ensuring dependencies are up to date...)
	dep ensure -v

schema: $(info $(LOG) Generating schemas...)
	# go generate ./schema
	# go-bindata -ignore=\.go -pkg=schema -o=schema/bindata.go schema/...
	go generate ./schema

wire: $(info $(LOG) Genrating services for DI...)
	wire ./service/

dev: dep; schema; $(info $(LOG) Running as developer...)
	go run ./main.go

build: dep ; $(info $(LOG) Golang is building in `go-spacex-graphql/dist`)
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../dist/main ./main.go

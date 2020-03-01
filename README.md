# go-spacex-graphql
:rocket: Go's GraphQL server to serve SpaceX API

## Requirements

- Golang
- Dep

### Recommended

- Make (Optional)
- Postman (Optional)

## Running Locally

### Clone the project using Git

```bash
git clone https://github.com/estebanborai/go-spacex-graphql.git
```

### Install project dependencies

```bash
cd /go-spacex-graphql/src/ && make dep
```

### Install utilities

#### go-bindata

In order to generate GraphQL schemas a package called *go-bindata* is required.
Install *go-bindata* using Golang's `get` command as follows:

```bash
go get -u github.com/jteeuwen/go-bindata/...
```

#### wire

*google/wire* is used to generate/build service providers. Install *google/wire* using the
following command:

```bash
go get github.com/google/wire/cmd/wire
```

### Run the project

```bash
# @ go-spacex-graphql/src/
go run ./main.go
```

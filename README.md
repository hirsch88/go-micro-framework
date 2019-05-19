# GoMicroFramework

## ❯ Why

Our main goal with this project is a feature complete server application.
We like you to be focused on your business and not spending hours in project configuration.

Try it!! We are happy to hear your feedback or any kind of new features.

## ❯ Installation

```shell
brew install go
brew install dep
brew install goose
brew install mysql
```

## ❯ Getting Started

To start the server run.
```shell
go run main.go
```

To run the api test run.
```shell
go test -v ./tests
```

## ❯ Dependency Management

Our dependency management tool is [dep](https://golang.github.io/dep/);

### Add Dependency

```shell
dep ensure -add github.com/foo/bar
```

### Updating Dependencies

```shell
// dry run testing an update
dep ensure -update -n
// non-dry run
dep ensure -update
// updates a specific package
dep ensure -update github.com/gorilla/mux
// updates to a specific version
dep ensure -update github.com/gorilla/mux@1.0.0
```

## ❯ Database Migrations

To manage our database schema we used the library [goose](https://github.com/pressly/goose).

### Run migrations

```shell
goose -dir database/migrations mysql "root:root@/go-micro-framework?parseTime=true" up 
```

### Create a migration

```shell
goose -dir database/migrations create create_user_table sql
```

## ❯ Testing

For testing our awesome app we decided to use the toolkit [stretchr/testify](https://github.com/stretchr/testify).

### Create Mocks

To automatically generate the mocks we use [vektra/mockery](https://github.com/vektra/mockery). 
The below command example generates the mock files of our interfaces inside the /app folder.

```shell
mockery -dir=./app -recursive=true -all 
```

## ❯ License

[MIT](/LICENSE)

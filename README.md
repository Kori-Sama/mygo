## This is a backend server for MyGO!!!!! project

### Setup

Please download go 1.21.6

then

```shell
go mod download
```

finally

```shell
go run ./cmd
```

maybe you can run this project, but I'm not sure🧐

### Swagger

```shell
swag init -g cmd/main.go -o docs
```

run the project and open http://127.0.0.1:8888/swagger/index.html
and then you can see the swagger docs page


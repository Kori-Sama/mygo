## This is a backend server for MyGO!!!!! project

### Front-end

[MyGO-GUI](https://github.com/MyGO-GUI-Project/MyGO-GUI)

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

### Database

execute the sql file named init.sql in the scripts folder
it will create schema and tables automatically

### Swagger

```shell
swag init -g cmd/main.go -o docs
```

run the project and open http://127.0.0.1:8888/swagger/index.html
and then you can see the swagger docs page


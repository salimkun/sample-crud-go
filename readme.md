### Simple Mono Repo With Golang

#Detail

- This app just for register user (1 API for register)

#How to use

**without docker**
- Open terminal and run :
```console
foo@bar:~$ cd sample-crud-go
foo@bar/sample-crud-go:~$ go get
foo@bar/sample-crud-go:~$ go mod tidy
foo@bar/sample-crud-go:~$ go run main.go
```
- Open Postman and Import postman collection from directory doc

**with docker**
- Open terminal and run :
```console
foo@bar:~$ cd sample-crud-go
foo@bar/sample-crud-go:~$ docker-compose build
foo@bar/sample-crud-go:~$ docker-compose up
```
- Open Postman and Import postman collection from directory doc
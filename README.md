# Coding Test

> Task: Create a server application, which can calculate simple math tasks. Also, implement a 
> client which gets input via program arguments and send it to the server. After the server 
> response, show the result on STDOUT.

```sh
$ ./client -method add -a 2 -b 3
$ 5
$ ./client -method sub -a 4 -b 2
$ 2
```

## Project Structure

This project is structured according to [Standard Go Project Layout](https://github.com/golang-standards/project-layout#go-directories) in the following way:

* [`api`](./api): Protobuf definition files to specify the communication protocol between client and server
* [`cmd`](./cmd): Main applications for this project
* [`internal`](./internal): Private application and library code of this project
* [`scripts`](./scripts): Scripts to perform various operations

## Quickstart

Install the latest version of Go. For more information, visit [go.dev/doc/install](https://go.dev/doc/install).

Build both client and server applications by invoking `make build` (if [make](https://www.gnu.org/software/make/) is installed on your system) or otherwise manually using 

```sh
$ go build github.com/felix-kaestner/calculator/cmd/server
$ go build github.com/felix-kaestner/calculator/cmd/client
```

Start the server application

```sh
$ ./server
2022/03/11 10:45:00 Server listening at [::]:8000
```

Afterwards, open a new terminal and use the client

```sh
$ ./client -method add -a 2 -b 3
$ 5
```

## System Requirements

This project only uses the Go language and dependencies written entirely in Go. 

Thus, the [Go System Requirements](https://doc.codingdict.com/golang/doc/golang.org/doc/install.html#requirements) are sufficient.

For development, please follow the system requirements linked in the [`api`](./api) directory. Additionally, [make](https://www.gnu.org/software/make/) can aid the development process by defining common tasks inside the [Makefile](./Makefile).

## Dependencies

Looking at the [go.mod](./go.mod) file this project only has two main direct dependencies:

* `google.golang.org/grpc`: The gRPC library
* `google.golang.org/protobuf`: The protobuf library

## Cheers ✌
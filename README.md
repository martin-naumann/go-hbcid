# Go-HBCI-Server

## Description
A simplistic HBCI server written in GO.
It allows to log in via PIN/TAN in FinTS 3.00 and do some basic jobs:

* Authenticate a user
* Get the accounts of a user (not implemented yet)
* Get the balance of each account (not implemented yet)
* Get the transactions of each account (not implemented yet)
* Simulate a transfer (not implemented yet)

## Architecture

### What it does
* main.go starts an HTTPS server on port 8080 with a self-signed certificate
* It accepts Base64 encoded HTTP POST requests (as usual for PIN/TAN)
* The payload is decoded and the FinTS messages are parsed
* Depending on the parsed payload a response is generated and written back as the response

### Structure

```shell
    ├── Makefile
    ├── README.md
    ├── Vagrantfile
    ├── provisioning
    │   ├── setup_riak.sh
    │   └── setup_server.sh
    ├── server.crt
    ├── server.key
    └── src
        ├── hbci
        │   ├── message.go
        │   ├── message_test.go
        │   └── types.go
        └── main.go
```

* `src/main.go`: The entrypoint of the code. Here you find the HTTPS server code, the request handlers and the parsing of messages
* `src/hbci/message.go`: Code to generate messages and unwrap encryption headers
* `src/hbci/types.go`: Some data structures used for HBCI
* `server.key` / `server.key`: SSL Certificate and key for HTTPS
* `provisioning`: Contains shell scripts for the initial provisioning of the vagrant boxes

## Dependencies
This project requires
* Go v1.2+
* git
* Redis for persistence
* (Go dependency) testify for assertions in the tests
* (Go dependency) glog for leveled logging
* (Go dependency) lint for linting on `make lint`
* (Go dependnency) uniuri for generating random strings
* (Go dependency) redis for talking to the redis server
* fpm (Ruby 1.9+) is used to create a `deb` package when running `make package`

## Getting Started

### Bringing up vagrant
It is recommended to use the vagrant boxes to run this:
```shell
  vagrant up
```

It will bring up a virtual machine called "go-hbcid-server" with the IP 10.10.10.2 where you can run, test and build the server.

### Adding users

Make a POST request with `login_id` and `pin` to `https://10.10.10.2/users` like this:

```shell
  $  curl -k https://10.10.10.2:8080/users -X POST -d "login_id=user1&pin=123def"
```
The call above will create the user `user1` with the PIN `123def`.

### Talking to the HBCI server

The server accepts all BLZs (it doesn't care at all), HBCI 3.00 and uses the URL https://10.10.10.2:8080/hbci

### Running the server with a different loglevel
If you want to run the HBCI server with a different log level, append `-stderrthreshold=<Minimum log level>`, e.g.
```shell
   $ sudo build/server -stderrthreshold=WARNING
```

## Deploy
The server can be deployed as a debian package or directly using the binary.

## Notes

### Running the server on a different port
By default the server starts to listen on all addresses on port 8080.
However, that can be adjusted by passing `--addr` on startup, like this:

```shell
    $ sudo build/server --addr :443
```

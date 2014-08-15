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

* `src/main.go`: The entrypoint of the code. Here you find the HTTPS server code, the response handler and the parsing of the message
* `src/hbci/message.go`: Code to generate messages and unwrap encryption headers
* `src/hbci/types.go`: Some data structures used for HBCI
* `server.key` / `server.key`: SSL Certificate and key for HTTPS
* `provisioning`: Contains shell scripts for the initial provisioning of the vagrant boxes

## Dependencies
This project requires
* Go v1.2+
* git
* (Go dependency) testify for assertions in the tests
* (Go dependency) glog for leveled logging
* (Go dependency) lint for linting on `make lint`
* (Go dependnency) uniuri for generating random strings
* fpm (Ruby 1.9+) is used to create a `deb` package when running `make package`

## Getting Started

### Bringing up vagrant
It is recommended to use the vagrant boxes to run this:
```shell
  vagrant up
```

It will bring up two virtual machines:
* go-hbcid-riak (10.10.10.20): A Riak instance for persistence
* go-hbcid-server (10.10.10.2): The VM for developing, building and running the HBCI server in

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

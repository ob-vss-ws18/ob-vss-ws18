# Greeter Service

This is the Greeter service

Generated with

```
micro new github.com/ob-vss-ws18/ob-vss-ws18/micro/greeter --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.greeter
- Type: srv
- Alias: greeter

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./greeter-srv
```

Build a docker image
```
make docker
```
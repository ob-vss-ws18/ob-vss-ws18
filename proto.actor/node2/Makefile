PLATFORM := $(shell uname -s)

app: main.go
ifeq ($(PLATFORM), Linux)
	go get golang.org/x/sys/unix
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o node2_docker main.go
else
	go build -o node2 main.go
endif

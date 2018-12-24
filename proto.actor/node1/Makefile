PLATFORM := $(shell uname -s)

app:
ifeq ($(PLATFORM), Linux)
	go get golang.org/x/sys/unix
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o node1_docker main.go
else
	go build -o node1 main.go
endif

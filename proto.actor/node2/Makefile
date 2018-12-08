PLATFORM := $(shell uname -s)

app: main.go
ifeq ($(PLATFORM), Linux)
	go get golang.org/x/sys/unix
	go build -o node2_docker main.go
else
	go build -o node2 main.go
endif

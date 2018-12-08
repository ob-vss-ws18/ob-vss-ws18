PLATFORM := $(shell uname -s)

app:
ifeq ($(PLATFORM), Linux)
	go get golang.org/x/sys/unix
	go build -o node1_docker main.go
else
	go build -o node1 main.go
endif

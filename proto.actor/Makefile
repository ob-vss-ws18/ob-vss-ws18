
build:
	cd node1 && make
	cd node2 && make

docker-build:
	docker run --rm -v ~/go/src:/go/src \
		-w /go/src/github.com/ob-vss-ws18/ob-vss-ws18/proto.actor/node1 iron/go:dev make app
	# cd node1 && docker build -t node1 .
	# rm node1/node1_docker
	docker run --rm -v ~/go/src:/go/src \
		-w /go/src/github.com/ob-vss-ws18/ob-vss-ws18/proto.actor/node2 iron/go:dev make app
	# cd node2 && docker build -t node2 .
	# rm node2/node2_docker

	
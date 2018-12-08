FROM iron/go

RUN addgroup -S app && adduser -S -G app app 
USER app

EXPOSE 8090

WORKDIR /app
ADD node1_docker /app/
ENTRYPOINT ["/app/node1_docker", "--bind=node1:8090", "--node2host=node2:8091"]
FROM iron/go

RUN addgroup -S app && adduser -S -G app app 
USER app

EXPOSE 8091

WORKDIR /app
ADD node2_docker /app/
ENTRYPOINT ["/app/node2_docker", "--bind=node2:8091"]
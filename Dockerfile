FROM golang:1.23

WORKDIR /app

COPY . .

ENV MONGODB_HOST="host.docker.internal"

RUN make build

CMD ["bin/main"]
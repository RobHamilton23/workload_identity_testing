# Build command:
# docker buildx build \
#   --builder=container \
#   --platform=linux/amd64 \
#   --push \
#   -t quay.io/robhamilton/workload_identity_testing:0.0.8 .

FROM golang:1.20.3-alpine3.16

WORKDIR /app

COPY main.go /app
COPY go.mod /app
COPY go.sum /app

RUN go build .
CMD /app/workload_identity_testing

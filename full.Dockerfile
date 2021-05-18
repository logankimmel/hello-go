FROM golang:1.15
COPY src/main.go /go/src/hello-go/main.go
WORKDIR /go/src/hello-go
RUN go get ./... && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /hello-go .
CMD ["/hello-go"]
EXPOSE 8080

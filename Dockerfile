FROM golang:1.15 as build
COPY main.go /go/src/hello-go/main.go
WORKDIR /go/src/hello-go
RUN go get ./... && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /hello-go .

FROM scratch
COPY --from=build /hello-go /hello-go
CMD ["/hello-go"]
EXPOSE 8080

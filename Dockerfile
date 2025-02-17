FROM golang:1 as builder

COPY . /go/src/fdnBaseAPI

WORKDIR /go/src/fdnBaseAPI
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/fdnBaseAPI/ .
CMD ["./app"]

FROM golang:1.12 as builder
MAINTAINER niloofar.gheibi@zalando.de

WORKDIR /go/src/hellodocker

COPY main.go .
RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
EXPOSE 60234
WORKDIR /root/
COPY --from=builder /go/src/hellodocker .
CMD ["./app"]




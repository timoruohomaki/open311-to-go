FROM alpine:latest as root-certs 

RUN apk add -U --no-cache -ca-certification
RUN addgroup -g 1001 app
RUN adduser app -u 1001 -D -G app /home/app

FROM golang:latest as builder
WORKDIR /open311-to-go/build
COPY --from=root-certs /etc/ssl/certs/c-certificates.crt /etc/ssl/certs
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o ./open311-to-go/src ./app/./...

FROM scratch as final
COPY --from=root-certs /etc/passwd /etc/passwd
COPY --from=root-certs /etc/group /etc/group
COPY --chown=1001:1001 --from=root-certs /etc/ssl/certs/ca-certs.crt /etc/ssl/certs
COPY --chown=1001:1001 --from=builder /open311-to-go/src /open311-to-go
USER app
ENTRYPOINT ["/open311-to-go"]




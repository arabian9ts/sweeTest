FROM golang:latest as builder

ENV CGO_ENABLED=0
WORKDIR /go/src/github.com/arabian9ts/sweeTest
COPY . .
RUN make setup
RUN SWEETEST_ENV=container GOOS=linux GOARCH=amd64 go build -o sweeTest

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/arabian9ts/sweeTest/sweeTest sweeTest
COPY --from=builder /go/src/github.com/arabian9ts/sweeTest/config/envs config/envs
COPY --from=builder /go/src/github.com/arabian9ts/sweeTest/db db
COPY --from=builder /go/src/github.com/arabian9ts/sweeTest/bootstrap.sh bootstrap.sh
EXPOSE 8080

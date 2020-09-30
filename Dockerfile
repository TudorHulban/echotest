FROM golang:1.14 as build

WORKDIR /echotest
COPY go.mod ./
RUN GOSUMDB=off go mod download

COPY . .
RUN GOSUMDB=off go build -o echotest ./cmd/main.go

FROM debian:stretch-slim
USER nobody

WORKDIR /app
COPY --chown=nobody:users --from=build /echotest /app/echotest 

EXPOSE 1323/tcp

CMD ["/app/echotest/echotest"]
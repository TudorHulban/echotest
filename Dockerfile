FROM golang:1.14 as build

WORKDIR /echotest
COPY go.mod ./
RUN GOSUMDB=off go mod download

COPY . .
RUN GOSUMDB=off go build -o echotest ./cmd/main.go

FROM golang:1.14

WORKDIR /app
COPY --from=build /echotest /app/echotest 
EXPOSE 1323/tcp

CMD ["./echotest"]
FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server cmd/microservice/main.go

#FROM scratch 
FROM alpine:latest

COPY --from=builder /app/server .


CMD [ "./server" ]

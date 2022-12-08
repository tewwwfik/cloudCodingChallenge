FROM docker.io/golang:alpine as builder

RUN mkdir /app
ADD main.go /app/main.go
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

FROM scratch

WORKDIR /app/
COPY --from=builder /app/main /app/main

EXPOSE 8080
CMD ["./main"]

# app/Dockerfile
FROM golang:1.22.5

WORKDIR /app

COPY . .

RUN go build -o /app/app .

CMD ["/app/app"]

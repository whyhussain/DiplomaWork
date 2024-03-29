FROM golang:1.19.9-bullseye

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./built

RUN chmod +x built

EXPOSE 8080

CMD ["./built"]

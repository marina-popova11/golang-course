FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /collector ./cmd/collector/main.go

EXPOSE 8081

CMD ["/collector"]
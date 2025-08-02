FROM golang:1.24

WORKDIR /app

RUN apt-get update && apt-get install -y vim curl git

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]

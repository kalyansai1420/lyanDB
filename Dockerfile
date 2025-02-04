FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
    build-essential \
    golang \
    libc6-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o lyanDB main.go resp.go commands.go server.go storage.go

EXPOSE 6379

CMD ["./lyanDB"]

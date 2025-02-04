
# lyanDB - A Simple In-Memory Database

lyanDB is a lightweight, in-memory database. It's designed to be a basic key-value store that supports core Redis-like functionality. This project was created to explore various concepts including multithreading, network programming, and system design.

## Features

- **Basic Redis Commands**: 
  - `SET`, `GET`, `DEL`, `EXISTS`
  - `INCR`, `DECR`, `LPUSH`, `RPUSH`
  - `PING`, `ECHO`
  - Support for expiry options (`EX`, `PX`, `EXAT`, `PXAT`)
  
- **Multithreading**: 
  - Built using Go's goroutines to handle concurrent connections, making it scalable.
  
- **Persistence**: 
  - Save the database to a Redis Database (RDB) file using the `SAVE` command.
  
- **Dockerized**: 
  - Easily deployable using Docker. Just pull the image from Docker Hub and start the server.

- **Testing with `redis-benchmark`**: 
  - Use `redis-benchmark` to test the performance and scalability of the database.

## Getting Started

### Prerequisites

- Docker installed on your machine.

### Running lyanDB with Docker

To run `lyanDB` in a Docker container, use the following command:

```bash
docker run -p 6379:6379 saikalyan1420/lyandb
```

This will start the lyanDB server on port `6379`, which is the default Redis port.

### Build lyanDB Docker Image

If you want to build the Docker image locally, clone the repository and run:

```bash
docker build -t saikalyan1420/lyandb .
```

Then, run the image with:

```bash
docker run -p 6379:6379 saikalyan1420/lyandb
```

## Using lyanDB

Once the server is running, you can interact with it using any Redis-compatible CLI tool such as `redis-cli` or `telnet`.

### Example with `redis-cli`

- **SET Command**:

```bash
redis-cli -h localhost -p 6379 SET mykey "hello"
```

- **GET Command**:

```bash
redis-cli -h localhost -p 6379 GET mykey
```

- **DEL Command**:

```bash
redis-cli -h localhost -p 6379 DEL mykey
```

### Example with `telnet`

1. Open a terminal and run:

```bash
telnet localhost 6379
```

2. Type the following commands:

```bash
SET mykey "hello"
GET mykey
DEL mykey
```

## Performance Benchmarking

You can use the Redis benchmarking tool `redis-benchmark` to test the performance of the database. Here's how you can run it:

```bash
redis-benchmark -t set,get -n 100000 -q
```

This will test the `SET` and `GET` commands with 100,000 requests and provide you with performance metrics.


## Contributing

Feel free to open issues, fork the project, and submit pull requests. Contributions to enhance functionality, improve performance, or fix bugs are welcome.

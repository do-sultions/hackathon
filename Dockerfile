FROM golang:1.18
WORKDIR /app
COPY . /app
CMD ["/usr/local/go/bin/go", "run", "main.go"]
EXPOSE 8080

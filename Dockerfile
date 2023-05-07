FROM golang:latest

WORKDIR /app

ENTRYPOINT ["tail", "-f", "/dev/null"]
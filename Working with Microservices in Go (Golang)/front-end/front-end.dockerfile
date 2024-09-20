FROM golang:1.23-alpine

RUN mkdir /app

COPY . /app

# Run the server executable
CMD [ "go run", "/app/cmd/web" ]
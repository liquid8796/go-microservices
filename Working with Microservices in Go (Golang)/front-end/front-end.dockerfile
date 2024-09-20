FROM alpine:latest

RUN mkdir /app

COPY . /app

# Run the server executable
CMD [ "go run", "/app/cmd/web" ]
FROM alpine:latest

RUN mkdir /app

COPY . /app

# Run the server executable
CMD [ "/app/cmd/web" ]
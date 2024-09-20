FROM alpine:latest

RUN mkdir /app

COPY front-end /app

# Run the server executable
CMD [ "/app/frontEndLinux" ]
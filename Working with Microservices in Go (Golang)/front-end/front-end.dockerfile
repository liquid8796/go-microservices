FROM golang:1.23-alpine

RUN mkdir /app

COPY . /app

EXPOSE 8083

# Run the server executable
CMD [ "/app/cmd/web" ]
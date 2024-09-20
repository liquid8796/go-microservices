FROM golang:1.23-alpine

RUN mkdir /app

COPY . /app

RUN chmod +x /app/cmd/web

EXPOSE 8083

# Run the server executable
CMD [ "/app/cmd/web" ]
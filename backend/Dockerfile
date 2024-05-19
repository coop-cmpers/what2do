FROM golang:alpine

RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .
#COPY .env .

RUN chmod +x ./docker-entrypoint.sh ./wait-for-it.sh

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /build ./src/main.go

EXPOSE ${BACKEND_PORT}

ENTRYPOINT ["./docker-entrypoint.sh"]
CMD ["/build"]
FROM golang:alpine

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .
COPY .env .

RUN go get -u github.com/swaggo/swag/cmd/swag

RUN $GOPATH/bin/swag init

RUN go get -d -v ./...

# RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN go install -v ./...

RUN go build -o /build

EXPOSE 8080 8080

CMD [ "/build" ]
FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /rmsubekti

COPY . .

RUN go mod tidy

RUN go build -o golang-devcode-todo

ENTRYPOINT [ "/rmsubekti/golang-devcode-todo" ]

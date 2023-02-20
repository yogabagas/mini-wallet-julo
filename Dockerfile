FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go mod vendor

RUN go build -o binary

EXPOSE 9000

ENTRYPOINT ["/app/binary"]

CMD ["serve"]
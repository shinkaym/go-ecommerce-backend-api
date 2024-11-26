FROM golang:alpine

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8002

CMD ["go", "run", "./cmd/server"]
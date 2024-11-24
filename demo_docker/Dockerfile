# docker build . -t name-demo
# docker images
# docker run -p 8005:8080 name-demo

# FROM golang:1.23-alpine
# WORKDIR /build

# COPY . .

# RUN go build -o nameDemo .

# WORKDIR /dist
# RUN cp /build/nameDemo /dist/

# EXPOSE 8009

# CMD ["/dist/nameDemo"]


#-------------------------
  
# docker build . -t name-demo-1
# docker run -p 8004:8080 name-demo-1

FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY . .

RUN go build -o nameDemo1 .

FROM scratch

COPY --from=builder /build/nameDemo1 /

ENTRYPOINT ["/nameDemo1"]
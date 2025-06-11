FROM golang:1.24.2-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

RUN git clone https://github.com/lassejlv/go-http-serve.git .

RUN go mod tidy && go build -o server main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/server .

ENV PORT=3000
ENV STATIC_DIR=/static

VOLUME ["/static"]

EXPOSE 3000

CMD ./server -port ${PORT} -dir ${STATIC_DIR}

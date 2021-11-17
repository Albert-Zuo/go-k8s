# builder step
FROM golang:1.16-alpine as builder
LABEL maintainer="Albert-Zuo Developers <xiaozuo1221@gmail.com>" version="1.0"
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /app/monitor
CMD [ "/monitor" ]

# image step
FROM alpine:3.11
COPY --from=builder /app/monitor /monitor
CMD [ "/monitor" ]

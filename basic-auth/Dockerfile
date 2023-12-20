# Build Stage
FROM golang:1.21.0-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go get ./...
RUN go build -o basic_auth main.go

# Deploy Stage
FROM alpine:3.18.4
COPY . /app
COPY --from=builder /build/basic_auth /app/
WORKDIR /app
EXPOSE 3000
CMD [ "./basic_auth" ]

# EOF
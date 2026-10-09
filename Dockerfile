FROM golang:1.26-alpine AS builder
WORKDIR /workspace
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/goCartAPI ./cmd/api


FROM alpine:3.20
RUN addgroup -S app && adduser -S app -G app
WORKDIR /server
COPY --from=builder /build/goCartAPI /server/goCartAPI
USER app
EXPOSE 3000
ENTRYPOINT ["/server/goCartAPI"]









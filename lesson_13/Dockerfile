# ===== build stage =====
FROM golang:1.25-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /out/server ./lesson_13/cmd/server

# ===== run stage =====
FROM alpine:3.20

WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY --from=builder /out/server ./server

EXPOSE 8080
CMD ["./server"]

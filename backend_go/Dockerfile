# syntax=docker/dockerfile:1.6

# =============== Build stage ===============
FROM golang:1.24-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git ca-certificates tzdata

# Lock to local toolchain; go.mod 'go' directive must be <= 1.22
ENV GOTOOLCHAIN=local
COPY backend_go/go.mod ./
COPY backend_go/go.sum ./
RUN go mod download

COPY backend_go/. ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/server ./main.go
    

# =============== Runtime stage ===============
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Jakarta
COPY --from=builder /out/server /app/server

# Copy .env if provided at runtime via bind-mount/secret (optional)

EXPOSE 8000
USER 65532:65532
ENTRYPOINT ["/app/server"]

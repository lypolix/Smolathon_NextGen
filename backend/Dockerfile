# ===== build stage
FROM golang:1.24 AS builder
WORKDIR /app

ENV GOPROXY="https://proxy.golang.org,direct" \
    GONOSUMDB="" \
    GOINSECURE="" \
    GODEBUG="netdns=go" \
    GOFLAGS="-modcacherw"

# Кеш модулей
COPY go.mod go.sum ./
RUN set -eux; \
    for i in 1 2 3; do \
      go mod download && break || (echo "go mod download retry $i/3" && sleep 2); \
    done

# Код
COPY . .

# Сборка с подробным выводом
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -v -o server ./cmd/backend

# ===== runtime
FROM gcr.io/distroless/base-debian12 AS runtime
WORKDIR /app
COPY --from=builder /app/server /app/server
ENV PORT=8080
EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/app/server"]

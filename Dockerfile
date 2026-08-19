# Stage 1: Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/meshery-mcp-server main.go

# Stage 2: Final lightweight runtime container
FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app
COPY --from=builder /app/meshery-mcp-server /app/meshery-mcp-server

EXPOSE 8080

ENTRYPOINT ["/app/meshery-mcp-server"]
CMD ["-transport=sse", "-port=8080"]

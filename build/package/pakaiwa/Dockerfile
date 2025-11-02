# Stage 1: Build
FROM golang:1.25.3-alpine3.22 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -trimpath -o pakaiwa ./cmd/pakaiwa

# Stage 2: Runtime (tanpa git, lebih aman!)
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /build/pakaiwa /app
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/app"]

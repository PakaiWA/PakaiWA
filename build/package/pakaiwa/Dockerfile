FROM golang:1.25.6-trixie AS builder
WORKDIR /build

# install build tools dan librdkafka-dev
RUN apt-get update && apt-get install -y \
    git \
    ca-certificates \
    build-essential \
    librdkafka-dev \
 && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -x -ldflags="-w -s" -trimpath -o pakaiwa ./cmd/pakaiwa

FROM gcr.io/distroless/base-debian12:nonroot
COPY --from=builder /build/pakaiwa /app
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/app"]

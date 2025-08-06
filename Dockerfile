FROM golang:1.24.5-alpine

# Install dependencies
RUN apk add --no-cache git bash curl

# Install Air (pakai module path yang benar!)
RUN go install github.com/air-verse/air@v1.62.0

WORKDIR /app

# Copy dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# (Optional) Buat direktori tmp
RUN mkdir -p /app/tmp && chmod 777 /app/tmp

EXPOSE 8080

# Jalankan air langsung dari /go/bin
CMD ["sh", "-c", "/go/bin/air -c .air.toml"]

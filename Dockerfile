FROM golang:1.21.3-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o l0 cmd/l0/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /build/l0 ./
COPY --from=builder /build/config/ ./

EXPOSE 80

CMD ["./l0"]

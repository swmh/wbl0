# WB Tech: level # 0 (Golang)

# Installation

```bash
git clone https://github.com/swmh/wbl0
cd wbl0
make compose-build && compose-up
```

# Usage

## Get Order

```bash
curl http://localhost:8080/orders/{order_uid}
```

## Create Orders

```bash
NATS_ADDRESS=nats://localhost:4222/ NATS_STREAM=orders NATS_CONSUMER=orders-consumer go run cmd/pub/main.go -n 5 -s orders
```

# Benchmarks
Specs:
- OS: Ubuntu 22.04.3 LTS x86_64
- Kernel: 5.15.0-91-generic
- CPU: AMD Ryzen 9 7950X (16) @ 4.491GHz
- RAM: 8GB

Processed messages: 9k message per second

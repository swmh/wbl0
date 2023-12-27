-- name: GetOrder :one
SELECT id, value FROM orders
WHERE id = $1 LIMIT 1;

-- name: GetNOrders :many
SELECT id, value FROM orders
LIMIT $1;

-- name: CreateOrder :exec
INSERT INTO orders (
  id, value
) VALUES (
  $1, $2
);

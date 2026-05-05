CREATE TABLE orders (
    id bigserial PRIMARY KEY,
    user_id bigint REFERENCES users(id),
    total numeric(10, 2),
    status text,
    placed_at timestamptz DEFAULT now()
);

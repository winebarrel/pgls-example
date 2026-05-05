CREATE TABLE users (
    id bigint PRIMARY KEY,
    email varchar(255) NOT NULL,
    name text,
    created_at timestamptz DEFAULT now()
);

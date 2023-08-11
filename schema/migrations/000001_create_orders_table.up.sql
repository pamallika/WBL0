CREATE TABLE IF NOT EXISTS orders(
    id serial PRIMARY KEY,
    data jsonb NOT NULL,
    vendor integer NOT NULL,
    status integer NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
    );
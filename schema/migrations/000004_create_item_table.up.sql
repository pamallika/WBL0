CREATE TABLE IF NOT EXISTS items
(
    id           serial PRIMARY KEY,
    order_id     int REFERENCES orders (id) NOT NULL,
    chrt_id      int       DEFAULT NULL,
    track_number varchar   DEFAULT NULL,
    price        int       DEFAULT NULL,
    rid          varchar   DEFAULT NULL,
    name         varchar   DEFAULT NULL,
    sale         int       DEFAULT NULL,
    size         varchar   DEFAULT NULL,
    total_price  int       DEFAULT NULL,
    nm_id       int       DEFAULT NULL,
    brand        varchar   DEFAULT NULL,
    status       int       DEFAULT NULL,
    created_at   TIMESTAMP DEFAULT now(),
    updated_at   TIMESTAMP DEFAULT now()
);
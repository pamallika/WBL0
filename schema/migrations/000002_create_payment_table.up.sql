CREATE TABLE IF NOT EXISTS payments
(
    id            serial PRIMARY KEY,
    transaction   varchar   DEFAULT NULL,
    request_id    varchar   DEFAULT NULL,
    currency      varchar   DEFAULT NULL,
    provider      varchar   DEFAULT NULL,
    amount        int       DEFAULT NULL,
    payment_dt    int       DEFAULT NULL,
    bank          varchar   DEFAULT NULL,
    delivery_cost int       DEFAULT NULL,
    goods_total   int       DEFAULT NULL,
    custom_fee   int       DEFAULT NULL,
    created_at    TIMESTAMP DEFAULT now(),
    updated_at    TIMESTAMP DEFAULT now()
);
CREATE TABLE IF NOT EXISTS deliveries
(
    id         serial PRIMARY KEY,
    name       varchar   DEFAULT NULL,
    phone      varchar   DEFAULT NULL,
    zip        varchar   DEFAULT NULL,
    city       varchar   DEFAULT NULL,
    address    varchar   DEFAULT NULL,
    region     varchar   DEFAULT NULL,
    email      varchar   DEFAULT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
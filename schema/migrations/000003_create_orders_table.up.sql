CREATE TABLE IF NOT EXISTS orders(
    id serial PRIMARY KEY,
    uid varchar DEFAULT NULL,
    track_number varchar DEFAULT NULL,
    entry varchar DEFAULT NULL,
    delivery_id int REFERENCES deliveries (id) DEFAULT NULL,
    payment_id int REFERENCES payments (id) DEFAULT NULL,
    locate varchar DEFAULT NULL,
    internal_signature varchar DEFAULT NULL,
    customer_id varchar DEFAULT NULL,
    delivery_service varchar DEFAULT NULL,
    shard_key varchar DEFAULT NULL,
    sm_id int DEFAULT NULL,
    date_created TIMESTAMP DEFAULT NULL,
    oof_shard varchar DEFAULT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
    );
CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR NOT NULL UNIQUE,
    data jsonb NOT NULL
    CONSTRAINT order_id_not_null_or_empty CHECK ( data ? 'order_uid' AND data->>'order_uid' != '' )
);

CREATE UNIQUE INDEX orders_orderId_idxjin ON orders ((data -> 'order_uid'));
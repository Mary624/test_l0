CREATE TABLE IF NOT EXISTS orders
(
    order_uid        VARCHAR(100) PRIMARY KEY,
    info     JSONB NOT NULL
);

-- CREATE TABLE IF NOT EXISTS deliveries
-- (
--     id uuid PRIMARY KEY,
--     name TEXT NOT NULL,
--     phone TEXT NOT NULL,
--     zip TEXT NOT NULL,
--     city TEXT NOT NULL,
--     address TEXT NOT NULL,
--     region TEXT NOT NULL,
--     email TEXT NOT NULL
-- );

-- CREATE TABLE IF NOT EXISTS payments
-- (
--     transaction uuid PRIMARY KEY,
--     request_id TEXT,
--     currency TEXT NOT NULL,
--     provider TEXT NOT NULL,
--     amount INTEGER NOT NULL,
--     payment_dt ????? INTEGER NOT NULL,
--     goods_total INTEGER NOT NULL,
--     custom_fee INTEGER NOT NULL
-- );

-- CREATE TABLE IF NOT EXISTS items
-- (
--     chrt_id INTEGER PRIMARY KEY,
--     track_number TEXT NOT NULL,
--     price INTEGER NOT NULL,
--     RID uuid NOT NULL,
--     name TEXT NOT NULL,
--     sale INTEGER NOT NULL,
--     size TEXT NOT NULL,
--     total_price INTEGER NOT NULL,
--     nm_id INTEGER NOT NULL,
--     brand TEXT NOT NULL,
--     status INTEGER NOT NULL
-- );

-- CREATE TABLE IF NOT EXISTS orders
-- (
--     order_uid        uuid PRIMARY KEY,
--     track_number     TEXT NOT NULL,
--     entry TEXT NOT NULL,
--     delivery_id uuid NOT NULL,
--     payment_id uuid NOT NULL,
--     items TEXT NOT NULL,//
--     locale TEXT NOT NULL,
--     internal_signature TEXT,
--     customer_id TEXT NOT NULL,
--     delivery_service TEXT NOT NULL,
--     shardkey TEXT NOT NULL,
--     sm_id INTEGER NOT NULL,
--     date_created timestamp NOT NULL,
--     oof_shard TEXT NOT NULL
-- );

-- CREATE TABLE IF NOT EXISTS orders_items
-- (
--     id        uuid PRIMARY KEY,
--     track_number     TEXT NOT NULL,
--     track_number     TEXT NOT NULL
-- );
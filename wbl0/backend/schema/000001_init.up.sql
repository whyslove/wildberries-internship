CREATE TABLE IF NOT EXISTS purchase(
    purchase_uid VARCHAR(255) PRIMARY KEY,
    track_number VARCHAR(255) UNIQUE,
    entry VARCHAR(255),
    locale VARCHAR(255),
    internal_signature VARCHAR(255),
    customer_id VARCHAR(255),
    delivery_service VARCHAR(255),
    shardkey VARCHAR(255),
    sm_id INTEGER,
    date_created TIMESTAMP,
    oof_shard VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS delivery(
    id SERIAL PRIMARY KEY,
    purchase_uid VARCHAR(255) UNIQUE
        REFERENCES purchase(purchase_uid) ON DELETE CASCADE,
    name VARCHAR(255),
    phone VARCHAR(255),
    zip VARCHAR(255),
    city VARCHAR(255),
    adress VARCHAR(255),
    region VARCHAR(255),
    email VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS payment (
    money_transaction VARCHAR(255)
        REFERENCES purchase(purchase_uid) ON DELETE CASCADE, 
    request_id VARCHAR(255),
    currency VARCHAR(255),
    provider VARCHAR(255),
    amount INTEGER CONSTRAINT positive_amount CHECK (amount > 0),
    payment_dt INTEGER,
    bank VARCHAR(255),
    delivery_cost INTEGER CONSTRAINT positive_delivery CHECK (delivery_cost > 0),
    goods_total INTEGER CONSTRAINT positive_total CHECK (goods_total > 0),
    custom_fee INTEGER CONSTRAINT positive_fee CHECK (custom_fee >= 0)
);

CREATE TABLE IF NOT EXISTS item (
    chrt_id INTEGER PRIMARY KEY,
    track_number VARCHAR(255),
    price INTEGER CONSTRAINT positive_price CHECK (price > 0),
    rid VARCHAR(255),
    name VARCHAR(255),
    sale INTEGER CONSTRAINT positive_sale CHECK (sale > 0),
    size VARCHAR(255),
    total_p INTEGER CONSTRAINT positive_t_price CHECK (
        total_p > 0
        AND total_p <= price
    ),
    nm_id INTEGER,
    brand VARCHAR(255),
    code_status INTEGER
);

CREATE TABLE IF NOT EXISTS purchase_item (
    purchase_uid VARCHAR(255)
        REFERENCES purchase(purchase_uid) ON DELETE CASCADE,
    item_chrt_id INTEGER
        REFERENCES item(chrt_id) ON DELETE CASCADE,
    amount INTEGER DEFAULT 1,
    CONSTRAINT purchase_item_pk PRIMARY KEY (purchase_uid, item_chrt_id)
);


create or replace procedure add_purchase(
    -- purchase table
    purchase_uid VARCHAR(255),
    track_number VARCHAR(255), 
    entry VARCHAR(255),
    locale VARCHAR(255),
    internal_signature VARCHAR(255),
    customer_id VARCHAR(255),
    delivery_service VARCHAR(255),
    shardkey VARCHAR(255),
    sm_id INTEGER,
    date_created TIMESTAMP,
    oof_shard VARCHAR(255)
)
LANGUAGE plpgsql    
AS $$
BEGIN
    -- subtracting the amount from the sender's account
    INSERT INTO purchase  
    VALUES
        (purchase_uid, track_number, entry, locale, internal_signature, customer_id, 
        delivery_service, shardkey, sm_id, date_created, oof_shard);
END$$;


create or replace procedure add_delivery(
    purchase_uid VARCHAR(255),
    name VARCHAR(255),
    phone VARCHAR(255),
    zip VARCHAR(255),
    city VARCHAR(255),
    adress VARCHAR(255),
    region VARCHAR(255),
    email VARCHAR(255)
)
LANGUAGE plpgsql    
AS $$
BEGIN
    -- subtracting the amount from the sender's account
    INSERT INTO delivery
    VALUES 
        (DEFAULT, purchase_uid, name, phone, zip, city, 
        adress, region, email);
END$$;

create or replace procedure add_payment(
    money_transaction VARCHAR(255),
    request_id VARCHAR(255),
    currency VARCHAR(255),
    provider VARCHAR(255),
    amount INTEGER,
    payment_dt INTEGER,
    bank VARCHAR(255),
    delivery_cost INTEGER,
    goods_total INTEGER,
    custom_fee INTEGER
)
LANGUAGE plpgsql    
AS $$
BEGIN
    -- subtracting the amount from the sender's account
    INSERT INTO payment
    VALUES
        (money_transaction, request_id, currency, provider, amount, 
        payment_dt, bank, delivery_cost, goods_total, custom_fee);
END$$;

create or replace procedure add_item_to_purchase(
    chrt_id INTEGER,
    track_number VARCHAR(255),
    price INTEGER,
    rid VARCHAR(255),
    name VARCHAR(255),
    sale INTEGER,
    size VARCHAR(255),
    total_p INTEGER,
    nm_id INTEGER,
    brand VARCHAR(255),
    code_status INTEGER,
    -- not item table values
    purchase_uid VARCHAR(255)
)
LANGUAGE plpgsql    
AS $$
BEGIN
    INSERT INTO item
    VALUES 
        (chrt_id, track_number, price, rid, name, sale, size, total_p,
        nm_id, brand, code_status) ON CONFLICT DO NOTHING;

    INSERT INTO purchase_item
    VALUES
        (purchase_uid, chrt_id) 
        ON CONFLICT ON CONSTRAINT purchase_item_pk DO UPDATE
            SET amount = purchase_item.amount + 1;
END$$;


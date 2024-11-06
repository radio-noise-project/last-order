CREATE TABLE IF NOT EXISTS sisters(
    sister_id UUID PRIMARY KEY,
    name VARCHAR(256) UNIQUE NOT NULL,
    role INT2 NOT NULL,
    address text NOT NULL,
    port INT4 NOT NULL,
    description text,
    private_key text NOT NULL,
    public_key text NOT NULL
);

CREATE TABLE IF NOT EXISTS keys(
    key_id UUID PRIMARY KEY,
    private_key text,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

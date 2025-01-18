CREATE TABLE key_values (
    user_id UUID NOT NULL,
    app_name TEXT NOT NULL,
    key TEXT NOT NULL,
    value TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, app_name, key)
);

CREATE TABLE api_keys (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    app_name TEXT NOT NULL,
    key TEXT NOT NULL,
    name TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_used TIMESTAMP,
    UNIQUE(key)
);

-- +goose Up
-- +goose StatementBegin
CREATE TABLE apps (
    id SERIAL PRIMARY KEY,
    created_by_user_id UUID NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE key_values (
    app_id INTEGER NOT NULL,
    key TEXT NOT NULL,
    value TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    PRIMARY KEY (app_id, key),
    FOREIGN KEY (app_id) REFERENCES apps(id)
);

CREATE TABLE api_keys (
    id SERIAL PRIMARY KEY,
    app_id INTEGER NOT NULL,
    key TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_used TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (app_id) REFERENCES apps(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE api_keys;
DROP TABLE key_values;
DROP TABLE apps;

-- +goose StatementEnd

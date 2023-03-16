-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS carts (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    purchased_at timestamptz,
    deleted_at timestamptz
);

CREATE TABLE IF NOT EXISTS carts_products (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    cart_id INT NOT NULL CONSTRAINT carts_fk REFERENCES carts(id),
    sku INT NOT NULL,
    cnt INT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamptz
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS notifications (
    id uuid PRIMARY KEY,
    type varchar(50) NOT NULL,
    coin varchar(50) NOT NULL,
    channel varchar(50) NOT NULL,
    data JSONB DEFAULT '{}'::jsonb, -- Store versions as JSON array
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT NULL,
    deleted_at timestamp DEFAULT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notifications CASCADE;
-- +goose StatementEnd

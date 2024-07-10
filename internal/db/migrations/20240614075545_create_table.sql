-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS notifications (
    id uuid PRIMARY KEY,
    type varchar(10) NOT NULL,
    coin varchar(50) NOT NULL,
    channels text,
    versions JSONB DEFAULT '{}'::jsonb, -- Store versions as JSON array
    count integer DEFAULT 1,
    created_by varchar(255) DEFAULT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT NULL,
    deleted_at timestamp DEFAULT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS cards CASCADE;
-- +goose StatementEnd

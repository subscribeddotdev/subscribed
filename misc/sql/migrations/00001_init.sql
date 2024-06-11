-- +goose Up
-- +goose StatementBegin
CREATE TABLE organizations (
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    disabled_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE organizations;
-- +goose StatementEnd

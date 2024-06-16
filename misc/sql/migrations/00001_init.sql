-- +goose Up
-- +goose StatementBegin
CREATE TABLE organizations (
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    disabled_at TIMESTAMP
);

CREATE TABLE members (
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    email TEXT NOT NULL UNIQUE,
    login_provider_id TEXT NOT NULL UNIQUE,
    organization_id VARCHAR(26) NOT NULL,
    created_at TIMESTAMP NOT NULL,

    CONSTRAINT pk_member_belongs_to_an_org FOREIGN KEY (organization_id) REFERENCES organizations (id)
);

CREATE TYPE EnvType AS ENUM('production', 'development');

CREATE TABLE environments (
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    organization_id VARCHAR(26) NOT NULL,
    name TEXT NOT NULL,
    env_type EnvType NOT NULL DEFAULT 'development',
    created_at TIMESTAMP NOT NULL,
    archived_at TIMESTAMP,
    CONSTRAINT pk_env_belongs_to_an_org FOREIGN KEY (organization_id) REFERENCES organizations (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE members;
DROP TABLE environments;
DROP TABLE organizations;
DROP TYPE EnvType;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE organizations (
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    disabled_at TIMESTAMPTZ
);

CREATE TABLE members (
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    email TEXT NOT NULL UNIQUE,
    login_provider_id TEXT NOT NULL UNIQUE,
    organization_id VARCHAR(26) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,

    CONSTRAINT pk_member_belongs_to_an_org FOREIGN KEY (organization_id) REFERENCES organizations (id)
);

CREATE TYPE EnvType AS ENUM('production', 'development');

CREATE TABLE environments (
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    organization_id VARCHAR(26) NOT NULL,
    name TEXT NOT NULL,
    env_type EnvType NOT NULL DEFAULT 'development',
    created_at TIMESTAMPTZ NOT NULL,
    archived_at TIMESTAMPTZ,
    CONSTRAINT pk_env_belongs_to_an_org FOREIGN KEY (organization_id) REFERENCES organizations (id)
);

CREATE TABLE applications (
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    environment_id VARCHAR(26) NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,

    CONSTRAINT pk_app_belongs_to_an_env FOREIGN KEY (environment_id) REFERENCES environments (id)
);

CREATE TABLE endpoints (
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    application_id VARCHAR(26) NOT NULL,
    url TEXT NOT NULL,
    description TEXT,
    signing_secret TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT pk_endpoint_belongs_to_an_app FOREIGN KEY (application_id) REFERENCES applications (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE endpoints;
DROP TABLE applications;
DROP TABLE members;
DROP TABLE environments;
DROP TABLE organizations;
DROP TYPE EnvType;
-- +goose StatementEnd

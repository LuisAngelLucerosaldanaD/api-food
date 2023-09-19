
-- +migrate Up
CREATE TABLE IF NOT EXISTS cfg.category(
    id BIGSERIAL  NOT NULL PRIMARY KEY,
    name VARCHAR (100) NOT NULL,
    role_id BIGINT  NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS cfg.category;

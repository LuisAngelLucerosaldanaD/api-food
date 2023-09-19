
-- +migrate Up
CREATE TABLE IF NOT EXISTS cfg.ingredient(
    id BIGSERIAL  NOT NULL PRIMARY KEY,
    name VARCHAR (255) NOT NULL,
    id_product BIGINT  NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS cfg.ingredient;

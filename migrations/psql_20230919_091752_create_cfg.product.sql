
-- +migrate Up
CREATE TABLE IF NOT EXISTS cfg.product(
    id BIGSERIAL  NOT NULL PRIMARY KEY,
    name VARCHAR (100) NOT NULL,
    url_img VARCHAR (255) NOT NULL,
    time VARCHAR (10) NOT NULL,
    description VARCHAR (500) NOT NULL,
    category_id BIGINT  NOT NULL,
    calorie VARCHAR (20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS cfg.product;

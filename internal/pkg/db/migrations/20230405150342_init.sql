-- +goose Up
-- +goose StatementBegin
CREATE TABLE warehouses (
                       id SERIAL PRIMARY KEY NOT NULL,
                       name text NOT NULL,
                       city text NOT NULL,
                       square INT,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);

CREATE TABLE products (
                     id SERIAL PRIMARY KEY NOT NULL,
                     name text NOT NULL,
                     description TEXT,
                     price INT NOT NULL,
                     warehouse_id INT,
                     created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
DROP TABLE warehouses;
-- +goose StatementEnd

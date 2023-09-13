DROP TABLE IF EXISTS products;

CREATE TABLE products (
                          id         SERIAL PRIMARY KEY,
                          name       VARCHAR(128) NOT NULL,
                          quantity   INTEGER NOT NULL,
                          code_value VARCHAR(128) NOT NULL,
                          is_published BOOLEAN NOT NULL DEFAULT FALSE,
                          expiration  VARCHAR(128) NOT NULL,
                          price      DECIMAL(5,2) NOT NULL
);

INSERT INTO products (name, quantity, price, code_value, is_published, expiration)
VALUES ('Product 1', 10, 10.00, '123456789', true, '2020-01-01');
-- +goose Up
CREATE SEQUENCE account_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    CACHE 1;
CREATE TABLE account (
    acc_id INT NOT NULL DEFAULT nextval('account_seq'),
    acc_email VARCHAR(255) NOT NULL,
    acc_password VARCHAR(255) NOT NULL,
    acc_last_login TIMESTAMPTZ NULL,
    acc_created_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    acc_created_by INT NOT NULL,
    acc_updated_date TIMESTAMPTZ NULL,
    acc_updated_by INT NOT NULL,
    CONSTRAINT acc_id PRIMARY KEY(acc_id)
);

-- +goose Down
DROP SEQUENCE IF EXISTS account_seq
DROP TABLE IF EXISTS account;
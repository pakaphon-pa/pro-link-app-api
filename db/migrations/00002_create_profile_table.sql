-- +goose Up
CREATE SEQUENCE profile_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    CACHE 1;
CREATE TABLE profile (
    prf_id INT NOT NULL DEFAULT nextval('profile_seq'),
    acc_id INT NOT NULL,
    prf_first_name VARCHAR(255) NOT NULL,
    prf_last_name VARCHAR(255) NOT NULL,
    prf_about VARCHAR(255) NULL,
    prf_phone_number VARCHAR(255) NULL,
    prf_phone_type VARCHAR(255) NULL,
    prf_address VARCHAR(255) NULL,
    prf_birth_month VARCHAR(255) NULL,
    prf_birth_date VARCHAR(255) NULL,
    prf_created_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    prf_created_by INT NOT NULL,
    prf_updated_date TIMESTAMPTZ NULL,
    prf_updated_by INT NOT NULL,
    CONSTRAINT PK_profile PRIMARY KEY(prf_id),
    CONSTRAINT FK_account FOREIGN KEY (acc_id) REFERENCES account(acc_id)
);

CREATE SEQUENCE website_profile_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    CACHE 1;
CREATE TABLE website_profile (
    web_id INT NOT NULL DEFAULT nextval('website_profile_seq'),
    prf_id INT NOT NULL,
    web_name VARCHAR(255) NOT NULL,
    web_type VARCHAR(255) NOT NULL,
    web_created_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    web_created_by INT NOT NULL,
    web_updated_date TIMESTAMPTZ NULL,
    web_updated_by INT NOT NULL,
    CONSTRAINT PK_website PRIMARY KEY(web_id),
    CONSTRAINT FK_profile FOREIGN KEY (prf_id) REFERENCES profile(prf_id)
);

-- +goose Down
DROP SEQUENCE IF EXISTS profile_seq
DROP TABLE IF EXISTS profile;
DROP FOREIGN KEY FK_account;

DROP SEQUENCE IF EXISTS website_profile_seq
DROP TABLE IF EXISTS website_profile;
DROP FOREIGN KEY FK_profile;
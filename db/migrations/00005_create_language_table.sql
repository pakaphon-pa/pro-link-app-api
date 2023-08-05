-- +goose Up
CREATE SEQUENCE language_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    CACHE 1;
CREATE TABLE language (
    lan_id INT NOT NULL DEFAULT nextval('language_seq'),
    acc_id INT NOT NULL,
    lan_name VARCHAR(255) NOT NULL,
    lan_proficiency VARCHAR(255) NOT NULL,
    lan_created_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    lan_created_by INT NOT NULL,
    lan_updated_date TIMESTAMPTZ NULL,
    lan_updated_by INT NOT NULL,
    CONSTRAINT PK_language PRIMARY KEY(lan_id),
    CONSTRAINT FK_account FOREIGN KEY (acc_id) REFERENCES account(acc_id)
);

-- +goose Down
DROP SEQUENCE IF EXISTS skill_seq
DROP TABLE IF EXISTS skill;
DROP FOREIGN KEY FK_account;

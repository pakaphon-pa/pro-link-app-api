-- +goose Up
CREATE SEQUENCE skill_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    CACHE 1;
CREATE TABLE skill (
    skl_id INT NOT NULL DEFAULT nextval('skill_seq'),
    acc_id INT NOT NULL,
    skl_name VARCHAR(255) NOT NULL,
    skl_created_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    skl_created_by INT NOT NULL,
    skl_updated_date TIMESTAMPTZ NULL,
    skl_updated_by INT NOT NULL,
    CONSTRAINT PK_skill PRIMARY KEY(skl_id),
    CONSTRAINT FK_account FOREIGN KEY (acc_id) REFERENCES account(acc_id)
);

-- +goose Down
DROP SEQUENCE IF EXISTS skill_seq
DROP TABLE IF EXISTS skill;
DROP FOREIGN KEY FK_account;

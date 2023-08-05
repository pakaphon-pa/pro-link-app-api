-- +goose Up
CREATE SEQUENCE experience_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    CACHE 1;
CREATE TABLE experience (
    exp_id INT NOT NULL DEFAULT nextval('experience_seq'),
    acc_id INT NOT NULL,
    exp_title VARCHAR(255) NOT NULL,
    exp_employee_type VARCHAR(255) NOT NULL,
    exp_company VARCHAR(255) NOT NULL,
    exp_company_locaiton VARCHAR(255) NULL,
    exp_location_type VARCHAR(255) NULL,
    exp_industry VARCHAR(255) NULL,
    exp_description TEXT NULL,
    exp_start_year INT NOT NULL,
    exp_start_month INT NULL,
    exp_is_current BOOLEAN NOT NULL DEFAULT FALSE,
    exp_end_year INT NULL,
    exp_end_month INT NULL,
    exp_created_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    exp_created_by INT NOT NULL,
    exp_updated_date TIMESTAMPTZ NULL,
    exp_updated_by INT NOT NULL,
    CONSTRAINT PK_experience PRIMARY KEY(exp_id),
    CONSTRAINT FK_account FOREIGN KEY (acc_id) REFERENCES account(acc_id)
);

-- +goose Down
DROP SEQUENCE IF EXISTS experience_seq
DROP TABLE IF EXISTS experience;
DROP FOREIGN KEY FK_account;

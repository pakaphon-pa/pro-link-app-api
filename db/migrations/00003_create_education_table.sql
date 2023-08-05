-- +goose Up
CREATE SEQUENCE education_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    CACHE 1;
CREATE TABLE education (
    edu_id INT NOT NULL DEFAULT nextval('education_seq'),
    acc_id INT NOT NULL,
    edu_school VARCHAR(255) NOT NULL,
    edu_degree VARCHAR(255) NOT NULL,
    edu_field_of_study VARCHAR(255) NOT NULL,
    edu_grade VARCHAR(255) NULL,
    edu_description TEXT NULL,
    edu_start_year INT NOT NULL,
    edu_start_month INT NULL,
    edu_end_year INT NOT NULL,
    edu_end_month INT NULL,
    edu_created_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    edu_created_by INT NOT NULL,
    edu_updated_date TIMESTAMPTZ NULL,
    edu_updated_by INT NOT NULL,
    CONSTRAINT PK_education PRIMARY KEY(edu_id),
    CONSTRAINT FK_account FOREIGN KEY (acc_id) REFERENCES account(acc_id)
);

-- +goose Down
DROP SEQUENCE IF EXISTS education_seq
DROP TABLE IF EXISTS education;
DROP FOREIGN KEY FK_account;

CREATE TYPE type_number as  enum('home', 'mobile', 'work', 'other');

CREATE TABLE user_phones(
    id serial PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    type type_number NOT NULL,
    number VARCHAR(20) NOT NULL,
    is_primary BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
);


alter table users add COLUMN user_phones_id serial REFERENCES user_phones(id);
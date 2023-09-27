CREATE TYPE type_number AS ENUM('home', 'mobile', 'work', 'other');

CREATE TABLE user_phones(
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    type type_number NOT NULL,
    number VARCHAR(20) NOT NULL,
    is_primary BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    UNIQUE(user_id, type)
);

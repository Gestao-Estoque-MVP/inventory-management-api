CREATE TABLE image (
    id serial PRIMARY KEY,
    description VARCHAR(255),
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
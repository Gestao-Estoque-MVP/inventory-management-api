CREATE TABLE image (
    id CHAR(36) PRIMARY KEY,
    description VARCHAR(255),
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

ALTER TABLE users 
    ADD COLUMN image_id VARCHAR NOT NULL REFERENCES image(id);
CREATE TABLE image (
    id UUID PRIMARY KEY,
    description VARCHAR(255),
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

ALTER TABLE users 
    ADD COLUMN image_id VARCHAR NOT NULL REFERENCES image(id);
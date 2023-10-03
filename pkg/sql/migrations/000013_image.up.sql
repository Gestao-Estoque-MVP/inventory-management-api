CREATE TABLE image (
    id UUID PRIMARY KEY,
    description VARCHAR(255),
    url VARCHAR(255)  NULL,
    created_at TIMESTAMP  NULL,
    updated_at TIMESTAMP
);

ALTER TABLE users 
    ADD COLUMN image_id UUID NULL REFERENCES image(id);
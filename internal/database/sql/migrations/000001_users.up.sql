CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NULL,
    email VARCHAR(100) NOT NULL,
    document VARCHAR(20) NOT NULL,
    password CHAR(64),
    active BOOLEAN DEFAULT FALSE,
    register_token VARCHAR(255),
    token_expires_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);
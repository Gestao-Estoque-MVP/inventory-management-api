CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NULL,
    email VARCHAR(100) NOT NULL,
    document VARCHAR(20) NOT NULL,
    password CHAR(64),
    mobile_phone VARCHAR(20),
    active BOOLEAN DEFAULT FALSE,
    register_token VARCHAR(255),
    token_expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
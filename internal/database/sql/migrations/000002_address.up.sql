CREATE TABLE address (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  address VARCHAR(255),
  street VARCHAR(255),
  city VARCHAR(100),
  state VARCHAR(100),
  postal_code VARCHAR(20),
  country VARCHAR(100),
  number VARCHAR(20),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE users
ADD COLUMN address_id UUID;
CREATE INDEX idx_address_id ON users (address_id);
CREATE TYPE tenant_type AS ENUM ('supplier', 'customer', 'super_admin');

CREATE TABLE tenant (
  id UUID PRIMARY KEY,
  name VARCHAR(255),
  tax_code VARCHAR(255),
  type tenant_type,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE users 
    ADD COLUMN tenant_id UUID NOT NULL REFERENCES tenant(id);


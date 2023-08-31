CREATE TYPE tenant_type AS ENUM ('supplier', 'customer', 'super_admin');

CREATE TABLE tenant (
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255),
  tax_code VARCHAR(255),
  type tenant_type
);

ALTER TABLE users 
    ADD COLUMN tenant_id VARCHAR NOT NULL REFERENCES tenant(id);


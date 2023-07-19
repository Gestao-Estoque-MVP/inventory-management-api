CREATE TABLE tenant (
  id INT PRIMARY KEY,
  name VARCHAR(255)
);

ALTER TABLE users 
    ADD COLUMN tenant_id INT REFERENCES tenant(id);


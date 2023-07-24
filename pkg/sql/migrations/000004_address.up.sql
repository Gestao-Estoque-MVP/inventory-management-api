CREATE TABLE address (
  id SERIAL PRIMARY KEY,
  user_id CHAR(36) UNIQUE NOT NULL,
  address VARCHAR(255),
  street VARCHAR(255),
  city VARCHAR(100),
  state VARCHAR(100),
  postal_code VARCHAR(20),
  country VARCHAR(100),
  FOREIGN KEY (user_id) REFERENCES users(id),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP
);


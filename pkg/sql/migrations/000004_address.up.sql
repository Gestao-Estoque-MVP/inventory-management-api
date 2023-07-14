CREATE TABLE address (
  id SERIAL PRIMARY KEY,
  user_id INTEGER,
  address VARCHAR(255),
  number VARCHAR(255),
  street VARCHAR(255),
  city VARCHAR(255),
  state VARCHAR(255),
  postal_code VARCHAR(20),
  country VARCHAR(255),
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE address (
  id UUID NULL ,
  user_id UUID UNIQUE  NULL,
  address VARCHAR(255),
  street VARCHAR(255),
  city VARCHAR(100),
  state VARCHAR(100),
  postal_code VARCHAR(20),
  country VARCHAR(100),
  number VARCHAR(20),
  FOREIGN KEY (user_id) REFERENCES users(id),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP
);


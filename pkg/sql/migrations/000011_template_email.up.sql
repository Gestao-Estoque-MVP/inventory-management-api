CREATE TABLE template_email (
   id VARCHAR(255) PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   url VARCHAR(255) NOT NULL,
   description VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
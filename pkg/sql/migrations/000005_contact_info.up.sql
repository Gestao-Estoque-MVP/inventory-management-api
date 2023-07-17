CREATE TABLE contact_info 
( 
 id VARCHAR(255) PRIMARY KEY,  
 name VARCHAR(255) NOT NULL,
 email VARCHAR(255) NOT NULL,
 phone VARCHAR(255),
 created_at TIMESTAMP NOT NULL,
 UNIQUE(email, phone)
);
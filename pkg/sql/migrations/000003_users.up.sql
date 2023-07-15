CREATE TABLE users 
( 
 id VARCHAR(255) PRIMARY KEY,  
 name VARCHAR(255) NOT NULL,    
 email VARCHAR(255) NOT NULL,
 phone VARCHAR(255),    
 document_type VARCHAR(255),   
 document_number VARCHAR(255),
 password VARCHAR(255),
 status VARCHAR(255) NOT NULL,
 register_token VARCHAR(255),
 token_expires_at TIMESTAMP,
 created_at TIMESTAMP NOT NULL
); 




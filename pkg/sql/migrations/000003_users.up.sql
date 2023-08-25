CREATE TYPE user_status AS ENUM ('pre-users', 'active', 'inative');

CREATE TABLE users 
( 
 id CHAR(36) PRIMARY KEY,  
 name VARCHAR(255) NOT NULL,    
 email VARCHAR(100) NOT NULL,
 phone VARCHAR(20) UNIQUE,    
 document_type VARCHAR(50),   
 document_number VARCHAR(50),
 password CHAR(64),
 status user_status NOT NULL,
 register_token VARCHAR(255),
 token_expires_at TIMESTAMP,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP,
 UNIQUE(email)
); 





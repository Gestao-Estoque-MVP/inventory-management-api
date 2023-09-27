CREATE TYPE user_status AS ENUM ('pre-users', 'active', 'inative');

CREATE TABLE users 
( 
 id UUID PRIMARY KEY,  
 name VARCHAR(255) NULL,    
 email VARCHAR(100) NOT NULL,
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





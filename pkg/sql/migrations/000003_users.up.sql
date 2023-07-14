CREATE TABLE users 
( 
 id SERIAL  PRIMARY KEY,  
 name VARCHAR(255) NOT NULL,   
 lastname VARCHAR(255) NOT NULL,   
 email VARCHAR(255) NOT NULL,
 phone VARCHAR(255) NOT NULL,   
 document_type VARCHAR(255) NOT NULL,   
 document_number VARCHAR(255) NOT NULL,
 password VARCHAR(255) NOT NULL
); 




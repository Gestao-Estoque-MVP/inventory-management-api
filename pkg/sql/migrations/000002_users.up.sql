CREATE TABLE users 
( 
 id VARCHAR(255) PRIMARY KEY,  
 name VARCHAR(255) NOT NULL,   
 lastname VARCHAR(255) NOT NULL,   
 email INT NOT NULL,   
 document_type VARCHAR(255) NOT NULL,   
 document_number VARCHAR(255) NOT NULL,   
 address_id INT
); 

ALTER TABLE users ADD FOREIGN KEY(address_id) REFERENCES address (ID) ON DELETE CASCADE ON UPDATE NO ACTION;



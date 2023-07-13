CREATE TABLE address 
( 
 id  SERIAL PRIMARY KEY ,  
 zip_code VARCHAR(255) NOT NULL,   
 address VARCHAR(255) NOT NULL,   
 state VARCHAR(255) NOT NULL,   
 city VARCHAR(255) NOT NULL
); 

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

--  migrate -path pkg/sql/migrations -database "postgresql://postgres:admin@172.18.0.4:5432/postgres?sslmode=disable" -verbose up  
--  migrate -path pkg/sql/migrations -database "postgresql://postgres:admin@172.18.0.4:5432/postgres?sslmode=disable" -verbose down
--   migrate -path pkg/sql/migrations -database "postgresql://postgres:admin@172.18.0.4:5432/postgres?sslmode=disable" force 1
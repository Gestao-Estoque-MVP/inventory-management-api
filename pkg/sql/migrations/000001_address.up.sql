CREATE TABLE address 
( 
 id  SERIAL PRIMARY KEY ,  
 zip_code VARCHAR(255) NOT NULL,   
 address VARCHAR(255) NOT NULL,  
 street VARCHAR(255) NOT NULL, 
 state VARCHAR(255) NOT NULL,   
 city VARCHAR(255) NOT NULL
);

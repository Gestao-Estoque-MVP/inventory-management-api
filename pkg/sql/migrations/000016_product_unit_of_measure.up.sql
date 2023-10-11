CREATE TABLE product_unit_of_measure (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE product
    ADD COLUMN product_unit_of_measure_id UUID REFERENCES product_unit_of_measure(id);
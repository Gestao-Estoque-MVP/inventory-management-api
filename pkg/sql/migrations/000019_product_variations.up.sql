CREATE TABLE product_variations(
    id UUID PRIMARY KEY,
    product_id UUID REFERENCES product(id) ON DELETE CASCADE,
    image_id UUID REFERENCES image(id ) ON DELETE CASCADE,
    price DECIMAL(10,2) NOT NULL,
    promotion DECIMAL(10,2) ,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

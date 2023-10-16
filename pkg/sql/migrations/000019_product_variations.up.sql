CREATE TABLE product_variations(
    id UUID PRIMARY KEY,
    product_id UUID REFERENCES product(id) ON DELETE CASCADE,
    image_id UUID REFERENCES image(id ) ON DELETE CASCADE,
    price NUMERIC(10,2) NOT NULL,
    promotion NUMERIC(10,2) ,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);

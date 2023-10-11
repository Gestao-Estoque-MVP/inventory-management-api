CREATE TABLE product_variation_mappings(
    id UUID PRIMARY KEY,
    product_variation_id UUID REFERENCES product_variations(id),
    variation_item_id UUID REFERENCES variation_items(id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
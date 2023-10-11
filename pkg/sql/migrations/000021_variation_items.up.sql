CREATE TABLE variation_items(
    id UUID PRIMARY KEY,
    variation_category_id UUID REFERENCES variation_categories(id),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
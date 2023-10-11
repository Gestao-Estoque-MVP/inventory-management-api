CREATE TABLE products_categories(
    id UUID PRIMARY KEY,
    product_id UUID REFERENCES product(id),
    category_id UUID REFERENCES category(id)
);
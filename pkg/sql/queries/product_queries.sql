-- name: CreateCategory :one
INSERT INTO category
    (id, name, description, created_at)
    VALUES($1, $2, $3, $4)
    RETURNING id;


-- name: CreateProduct :one
INSERT INTO product
    (id, 
    name, 
    low_stock_threshold, 
    image_id, 
    price, 
    tenant_id, 
    promotion, 
    safety_stock_level, 
    reorder_point, 
    min_lot, 
    max_lot, 
    fsn_classification, 
    width, 
    height, 
    length, 
    weight, 
    product_unit_of_measure_id
    )
    VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
    RETURNING id;


-- name: CreateProductUnitsOfMeasure :one
INSERT INTO product_unit_of_measure
    (id, name, description, created_at)
    VALUES($1, $2, $3, $4)
    RETURNING id;

-- name: CreateProductsCategories :one
INSERT INTO products_categories
    (id, product_id, category_id)
    VALUES($1, $2, $3)
    RETURNING id;
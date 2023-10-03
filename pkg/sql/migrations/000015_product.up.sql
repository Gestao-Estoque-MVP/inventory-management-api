CREATE TABLE "product" (
  "id" UUID PRIMARY KEY,
  "name" varchar(255),
  "low_stock_threshold" integer,
  "product_unit_of_measure_id" varchar(255),
  "image_id" varchar(255),
  "price" DECIMAL(10,2),
  "tenant_id" UUID,
  "promotion" DECIMAL(10,2),
  "safety_stock_level" INT,
  "reorder_point" INT,
  "min_lot" INT,
  "max_lot" INT,
  "fsn_classification" varchar(50),
  "width" int(100),
  "height" int(100),
  "length" int(100),
  "weight" int(100)
);

ALTER TABLE products
ADD FOREIGN KEY (tenant_id) REFERENCES tenant(id)
ON DELETE CASCADE; 

ALTER TABLE products
ADD FOREIGN KEY (image_id) REFERENCES image(id)
ON DELETE CASCADE;
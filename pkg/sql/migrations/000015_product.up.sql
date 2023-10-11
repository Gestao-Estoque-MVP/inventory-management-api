CREATE TABLE product (
  id UUID PRIMARY KEY,
  name varchar(255) NOT NULL,
  low_stock_threshold INT ,
  image_id UUID,
  price DECIMAL(10,2) NOT NULL,
  tenant_id UUID NOT NULL,
  promotion DECIMAL(10,2),
  safety_stock_level INT,
  reorder_point INT,
  min_lot INT,
  max_lot INT,
  fsn_classification varchar(50),
  width INT,
  height INT,
  length INT,
  weight INT
);

ALTER TABLE product
ADD FOREIGN KEY (tenant_id) REFERENCES tenant(id)
ON DELETE CASCADE; 

ALTER TABLE product
ADD FOREIGN KEY (image_id) REFERENCES image(id)
ON DELETE CASCADE;
BEGIN;

CREATE TYPE product_category AS ENUM ('Clothing', 'Accessories', 'Footwear', 'Beverages');

CREATE TABLE IF NOT EXISTS products (
  id VARCHAR(26) PRIMARY KEY NOT NULL,
  name VARCHAR(30) NOT NULL,
  sku VARCHAR(30) NOT NULL,
  category product_category NOT NULL,
  image_url TEXT NOT NULL,
  notes VARCHAR(200) NOT NULL,
  price INT NOT NULL,
  stock INT NOT NULL,
  location VARCHAR(200),
  is_available BOOLEAN NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_products_name ON products(lower(name));
CREATE INDEX IF NOT EXISTS idx_products_sku ON products(sku);
CREATE INDEX IF NOT EXISTS idx_products_category ON products(category);
CREATE INDEX IF NOT EXISTS idx_products_is_available ON products(is_available);
CREATE INDEX IF NOT EXISTS idx_products_in_stock ON products(stock) WHERE stock > 0;
CREATE INDEX IF NOT EXISTS idx_products_not_in_stock ON products(stock) WHERE stock = 0;
CREATE INDEX IF NOT EXISTS idx_products_created_at_desc ON products(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_products_created_at_asc ON products(created_at ASC);
CREATE INDEX IF NOT EXISTS idx_products_price_desc ON products(price DESC);
CREATE INDEX IF NOT EXISTS idx_products_price_asc ON products(price ASC);

COMMIT;
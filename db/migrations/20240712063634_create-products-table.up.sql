BEGIN;

DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'product_category') THEN
    CREATE TYPE product_category AS ENUM ('Clothing', 'Accessories', 'Footwear', 'Beverages');
  END IF;
END $$;

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

COMMIT;
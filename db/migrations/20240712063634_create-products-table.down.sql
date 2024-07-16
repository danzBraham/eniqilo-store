BEGIN;

DROP INDEX IF EXISTS idx_products_name;
DROP INDEX IF EXISTS idx_products_sku;
DROP INDEX IF EXISTS idx_products_category;
DROP INDEX IF EXISTS idx_products_is_available;
DROP INDEX IF EXISTS idx_products_in_stock;
DROP INDEX IF EXISTS idx_products_not_in_stock;
DROP INDEX IF EXISTS idx_products_created_at_desc;
DROP INDEX IF EXISTS idx_products_created_at_asc;
DROP INDEX IF EXISTS idx_products_price_desc;
DROP INDEX IF EXISTS idx_products_price_asc;

DROP TABLE IF EXISTS products;

DROP TYPE IF EXISTS product_category;

COMMIT;
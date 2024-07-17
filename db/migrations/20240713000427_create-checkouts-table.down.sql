BEGIN;

DROP INDEX IF EXISTS idx_checkouts_checkout_history_id;
DROP INDEX IF EXISTS idx_checkouts_product_id;
DROP INDEX IF EXISTS idx_checkouts_quantity;

DROP TABLE IF EXISTS checkouts;

COMMIT;
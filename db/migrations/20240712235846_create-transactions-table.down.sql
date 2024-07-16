BEGIN;

DROP INDEX IF EXISTS idx_checkout_histories_customer_id;
DROP INDEX IF EXISTS idx_checkout_histories_created_at_desc;
DROP INDEX IF EXISTS idx_checkout_histories_created_at_asc;

DROP TABLE IF EXISTS checkout_histories;

COMMIT;
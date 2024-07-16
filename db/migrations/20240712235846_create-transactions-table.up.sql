BEGIN;

CREATE TABLE IF NOT EXISTS checkout_histories (
  id VARCHAR(26) PRIMARY KEY NOT NULL,
  customer_id VARCHAR(26) NOT NULL,
  total_price INT NOT NULL DEFAULT 0,
  paid INT NOT NULL DEFAULT 0,
  change INT NOT NULL DEFAULT 0,
  is_deleted BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  FOREIGN KEY (customer_id) REFERENCES users(id) ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE INDEX IF NOT EXISTS idx_checkout_histories_customer_id ON checkout_histories(customer_id);
CREATE INDEX IF NOT EXISTS idx_checkout_histories_created_at_desc ON checkout_histories(created_at DESC);
CREATE INDEX IF NOT EXISTS idx_checkout_histories_created_at_asc ON checkout_histories(created_at ASC);

COMMIT;
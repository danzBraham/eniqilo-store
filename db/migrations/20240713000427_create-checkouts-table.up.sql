BEGIN;

CREATE TABLE IF NOT EXISTS checkouts (
  id VARCHAR(26) PRIMARY KEY NOT NULL,
  checkout_history_id VARCHAR(26) NOT NULL,
  product_id VARCHAR(26) NOT NULL,
  quantity INT NOT NULL,
  FOREIGN KEY (checkout_history_id) REFERENCES checkout_histories(id) ON DELETE NO ACTION ON UPDATE NO ACTION,
  FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE NO ACTION ON UPDATE NO ACTION
);

CREATE INDEX IF NOT EXISTS idx_checkouts_checkout_history_id ON checkouts(checkout_history_id);
CREATE INDEX IF NOT EXISTS idx_checkouts_product_id ON checkouts(product_id);
CREATE INDEX IF NOT EXISTS idx_checkouts_quantity ON checkouts(quantity);

COMMIT;
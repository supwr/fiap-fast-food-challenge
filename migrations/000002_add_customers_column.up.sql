ALTER TABLE sc_fast_food.customers ADD COLUMN created_at TIMESTAMP NOT NULL;
ALTER TABLE sc_fast_food.customers ADD COLUMN updated_at TIMESTAMP NULL;
ALTER TABLE sc_fast_food.customers ADD COLUMN deleted_at TIMESTAMP NULL;

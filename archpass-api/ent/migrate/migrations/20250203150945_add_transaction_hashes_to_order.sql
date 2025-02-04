-- Modify "orders" table
ALTER TABLE "orders" ADD COLUMN "mint_transaction_hash" character varying NULL, ADD COLUMN "transfer_transaction_hash" character varying NULL;

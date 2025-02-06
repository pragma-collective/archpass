-- Modify "tickets" table
ALTER TABLE "tickets" ALTER COLUMN "mint_price" DROP NOT NULL, ADD COLUMN "price_in_cents" bigint NULL;

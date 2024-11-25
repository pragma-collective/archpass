-- Create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "wallet_address" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "users_wallet_address_key" to table: "users"
CREATE UNIQUE INDEX "users_wallet_address_key" ON "users" ("wallet_address");

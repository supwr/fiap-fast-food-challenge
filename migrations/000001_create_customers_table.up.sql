CREATE TABLE IF NOT EXISTS sc_fast_food.customers (
    "id" BIGSERIAL NOT NULL,
    "document" INT NOT NULL UNIQUE,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    CONSTRAINT "PK_Customers" PRIMARY KEY ("id")
);
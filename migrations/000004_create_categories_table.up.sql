CREATE TABLE IF NOT EXISTS sc_fast_food.categories (
    "id" BIGSERIAL NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NULL,
    "active" BOOLEAN NOT NULL DEFAULT TRUE,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NULL,
    "deleted_at" TIMESTAMP NULL,
    CONSTRAINT "PK_Categories" PRIMARY KEY ("id")
);
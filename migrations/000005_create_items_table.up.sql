CREATE TABLE IF NOT EXISTS sc_fast_food.items (
    "id" BIGSERIAL NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NULL,
    "price" DECIMAL(10,2) NOT NULL,
    "category_id" BIGINT NOT NULL,
    "active" BOOLEAN NOT NULL DEFAULT TRUE,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NULL,
    "deleted_at" TIMESTAMP NULL,
    CONSTRAINT "PK_Items" PRIMARY KEY ("id"),
    CONSTRAINT "FK_Category" FOREIGN KEY ("category_id") REFERENCES sc_fast_food.categories(id)
);
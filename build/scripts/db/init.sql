BEGIN;

CREATE TABLE IF NOT EXISTS "merchant" (
    "name" TEXT,
    "email" TEXT UNIQUE,
    "perc" FLOAT,
    PRIMARY KEY ("name")
);

COMMIT;
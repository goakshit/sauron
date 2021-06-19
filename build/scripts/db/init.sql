BEGIN;

CREATE TABLE IF NOT EXISTS "merchant" (
    "name" TEXT,
    "email" TEXT,
    "perc" FLOAT,
    CONSTRAINT "merchant_pkey" PRIMARY KEY ("name"),
    CONSTRAINT "merchant_email_unique" UNIQUE ("email")
);

CREATE TABLE IF NOT EXISTS "user" (
    "name" TEXT,
    "email" TEXT UNIQUE,
    "due_amount" FLOAT, 
    "credit_limit" FLOAT,
    CONSTRAINT "user_pkey" PRIMARY KEY ("name"),
    CONSTRAINT "user_email_unique" UNIQUE ("email")
);

COMMIT;
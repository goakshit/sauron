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

CREATE TABLE IF NOT EXISTS "transaction" (
    "id" SERIAL,
    "user_name" TEXT,
    "merchant_name" TEXT, 
    "merchant_perc" FLOAT,
    "amount" FLOAT,
    CONSTRAINT "txn_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "fk_user" FOREIGN KEY(user_name) REFERENCES "user"(name),
    CONSTRAINT "fk_merchant" FOREIGN KEY(merchant_name) REFERENCES "merchant"(name)
);

CREATE TABLE IF NOT EXISTS "payback" (
    "id" SERIAL,
    "user_name" TEXT,
    "amount" FLOAT,
    CONSTRAINT "payback_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "fk_payback_user" FOREIGN KEY(user_name) REFERENCES "user"(name)
);

COMMIT;
CREATE TABLE IF NOT EXISTS public.accounts (
    "id" bigserial NOT NULL,
    "owner" TEXT NOT NULL,
    "balance" INTEGER NOT NULL,
    "currency" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS public.entries (
    "id" bigserial NOT NULL,
    "account_id" bigint NOT NULL,
    "amount" INTEGER NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("account_id") REFERENCES "accounts" ("id")
);

-- CREATE TABLE IF NOT EXISTS public.transfer (
--     "id" INTEGER NOT NULL,
--     "from_account_id" INTEGER NOT NULL,
--     "to_account_id" INTEGER NOT NULL,
--     "created_at" TIMESTAMP NOT NULL,
--     "updated_at" TIMESTAMP NOT NULL,
--     PRIMARY KEY ("id"),
--     FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id"),
--     FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id")
-- );

-- CREATE INDEX "idx_transfer_from_account_id" ON "transfer" ("from_account_id");

-- CREATE INDEX "idx_transfer_to_account_id" ON "transfer" ("to_account_id");






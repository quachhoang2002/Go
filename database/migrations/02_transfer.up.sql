CREATE TABLE IF NOT EXISTS public.transfer (
    "id" bigserial NOT NULL,
    "from_account_id" bigint NOT NULL,
    "to_account_id" bigint NOT NULL,
    "amount" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id"),
    FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id"),
    FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id")
);

CREATE INDEX "idx_transfer_from_account_id" ON "transfer" ("from_account_id");
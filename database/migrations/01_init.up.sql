CREATE TABLE IF NOT EXISTS public.todos
(
    id          SERIAL                   PRIMARY KEY,
    name        VARCHAR(255)             NOT NULL CONSTRAINT todos_name_check CHECK (name <> ''::text),
    description TEXT                     NOT NULL DEFAULT '',
    deleted_at  TIMESTAMP WITH TIME ZONE NULL     DEFAULT NULL,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

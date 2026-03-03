-- +goose Up
-- Migrate auth tables from integer/bigint IDs to UUID primary keys (preserves existing data).
-- Prerequisite: 001_auth_tables.sql already applied on grape_auth.

-- Phase 1: Add uuid columns and backfill (while old id columns still exist)

-- users
ALTER TABLE public.users ADD COLUMN id_uuid uuid;
UPDATE public.users SET id_uuid = gen_random_uuid();
ALTER TABLE public.users ADD COLUMN created_user_id_uuid uuid;
ALTER TABLE public.users ADD COLUMN updated_user_id_uuid uuid;
UPDATE public.users u SET created_user_id_uuid = u2.id_uuid FROM public.users u2 WHERE u2.id = u.created_user_id;
UPDATE public.users u SET updated_user_id_uuid = u2.id_uuid FROM public.users u2 WHERE u2.id = u.updated_user_id;

-- auth_logs (user_id, created_by_id, updated_by_id reference users)
ALTER TABLE public.auth_logs ADD COLUMN id_uuid uuid;
ALTER TABLE public.auth_logs ADD COLUMN user_id_uuid uuid;
ALTER TABLE public.auth_logs ADD COLUMN created_by_id_uuid uuid;
ALTER TABLE public.auth_logs ADD COLUMN updated_by_id_uuid uuid;
UPDATE public.auth_logs SET id_uuid = gen_random_uuid();
UPDATE public.auth_logs a SET user_id_uuid = u.id_uuid FROM public.users u WHERE u.id = a.user_id;
UPDATE public.auth_logs a SET created_by_id_uuid = u.id_uuid FROM public.users u WHERE u.id = a.created_by_id;
UPDATE public.auth_logs a SET updated_by_id_uuid = u.id_uuid FROM public.users u WHERE u.id = a.updated_by_id;

-- refresh_tokens
ALTER TABLE public.refresh_tokens ADD COLUMN id_uuid uuid;
ALTER TABLE public.refresh_tokens ADD COLUMN user_id_uuid uuid;
UPDATE public.refresh_tokens SET id_uuid = gen_random_uuid();
UPDATE public.refresh_tokens r SET user_id_uuid = u.id_uuid FROM public.users u WHERE u.id = r.user_id;

-- Phase 2: Drop FKs and PKs, drop old columns, rename uuid columns, add PKs and FKs

ALTER TABLE public.refresh_tokens DROP CONSTRAINT IF EXISTS refresh_tokens_user_id_fkey;
ALTER TABLE public.refresh_tokens DROP CONSTRAINT IF EXISTS refresh_tokens_pkey;
ALTER TABLE public.auth_logs DROP CONSTRAINT IF EXISTS auth_logs_pkey;
ALTER TABLE public.users DROP CONSTRAINT IF EXISTS users_pkey;

ALTER TABLE public.users DROP COLUMN id;
ALTER TABLE public.users RENAME COLUMN id_uuid TO id;
ALTER TABLE public.users ADD PRIMARY KEY (id);
ALTER TABLE public.users DROP COLUMN created_user_id;
ALTER TABLE public.users DROP COLUMN updated_user_id;
ALTER TABLE public.users RENAME COLUMN created_user_id_uuid TO created_user_id;
ALTER TABLE public.users RENAME COLUMN updated_user_id_uuid TO updated_user_id;

ALTER TABLE public.auth_logs DROP COLUMN id;
ALTER TABLE public.auth_logs DROP COLUMN user_id;
ALTER TABLE public.auth_logs DROP COLUMN created_by_id;
ALTER TABLE public.auth_logs DROP COLUMN updated_by_id;
ALTER TABLE public.auth_logs RENAME COLUMN id_uuid TO id;
ALTER TABLE public.auth_logs RENAME COLUMN user_id_uuid TO user_id;
ALTER TABLE public.auth_logs RENAME COLUMN created_by_id_uuid TO created_by_id;
ALTER TABLE public.auth_logs RENAME COLUMN updated_by_id_uuid TO updated_by_id;
ALTER TABLE public.auth_logs ADD PRIMARY KEY (id);
ALTER TABLE public.auth_logs ADD CONSTRAINT auth_logs_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;

ALTER TABLE public.refresh_tokens DROP COLUMN id;
ALTER TABLE public.refresh_tokens DROP COLUMN user_id;
ALTER TABLE public.refresh_tokens RENAME COLUMN id_uuid TO id;
ALTER TABLE public.refresh_tokens RENAME COLUMN user_id_uuid TO user_id;
ALTER TABLE public.refresh_tokens ADD PRIMARY KEY (id);
ALTER TABLE public.refresh_tokens ADD CONSTRAINT refresh_tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;

-- Recreate indexes that referenced old columns (indexes on id are replaced by PK; user_id indexes)
CREATE INDEX IF NOT EXISTS idx_auth_logs_user_id ON public.auth_logs USING btree (user_id);
CREATE INDEX IF NOT EXISTS idx_refresh_tokens_user_id ON public.refresh_tokens USING btree (user_id);

-- +goose Down
-- Reverting UUID -> int is lossy (cannot recover original integers); leave Down no-op or recreate from 001.
-- To roll back you would need to restore from backup.

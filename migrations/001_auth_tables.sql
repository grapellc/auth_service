-- +goose Up
-- Auth-only schema for grape_auth DB: users, auth_logs, refresh_tokens.
CREATE TABLE public.users (
    id bigserial NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    created_user_id bigint,
    updated_user_id bigint,
    clerk_id text,
    email text NOT NULL,
    username text,
    first_name text,
    last_name text,
    phone_number text,
    avatar_url text,
    password_hash text,
    role text DEFAULT 'user'::text,
    is_phone_verified boolean DEFAULT false,
    is_email_verified boolean DEFAULT false,
    PRIMARY KEY (id)
);
CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);

CREATE TABLE public.auth_logs (
    id serial PRIMARY KEY,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    created_by_id integer,
    updated_by_id integer,
    user_id integer,
    identifier character varying(255),
    action character varying(255),
    status character varying(50),
    failure_reason text,
    ip_address character varying(45),
    user_agent text,
    meta_data jsonb
);
CREATE INDEX idx_auth_logs_user_id ON public.auth_logs USING btree (user_id);
CREATE INDEX idx_auth_logs_identifier ON public.auth_logs USING btree (identifier);
CREATE INDEX idx_auth_logs_action ON public.auth_logs USING btree (action);
CREATE INDEX idx_auth_logs_status ON public.auth_logs USING btree (status);
CREATE INDEX idx_auth_logs_deleted_at ON public.auth_logs USING btree (deleted_at);

CREATE TABLE public.refresh_tokens (
    id serial PRIMARY KEY,
    user_id integer NOT NULL,
    token character varying(255) NOT NULL,
    expires_at timestamp with time zone NOT NULL,
    revoked boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    CONSTRAINT refresh_tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX idx_refresh_tokens_token ON public.refresh_tokens USING btree (token);
CREATE INDEX idx_refresh_tokens_user_id ON public.refresh_tokens USING btree (user_id);

-- +goose Down
DROP TABLE IF EXISTS public.refresh_tokens;
DROP TABLE IF EXISTS public.auth_logs;
DROP TABLE IF EXISTS public.users;

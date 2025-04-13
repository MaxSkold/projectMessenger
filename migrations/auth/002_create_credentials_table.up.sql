-- Table: auth.credentials

CREATE TABLE IF NOT EXISTS auth.credentials
(
    user_id uuid NOT NULL,
    email text COLLATE pg_catalog."default" NOT NULL,
    phone_number character varying(15) COLLATE pg_catalog."default",
    passhash text COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT credentials_pkey PRIMARY KEY (user_id),
    CONSTRAINT credentials_email_key UNIQUE (email),
    CONSTRAINT chk_phone_number CHECK (phone_number::text ~ '^\+?[1-9]\d{1,14}$'::text)
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS auth.credentials
    OWNER to postgres;


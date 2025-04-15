-- FUNCTION: auth.update_updated_at_column()

CREATE OR REPLACE FUNCTION auth.update_updated_at_column()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE NOT LEAKPROOF
AS $BODY$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$BODY$;

ALTER FUNCTION auth.update_updated_at_column()
    OWNER TO postgres;
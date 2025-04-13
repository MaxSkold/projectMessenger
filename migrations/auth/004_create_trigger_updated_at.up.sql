-- Trigger: set_updated_at

CREATE OR REPLACE TRIGGER set_updated_at
    BEFORE UPDATE
    ON auth.credentials
    FOR EACH ROW
EXECUTE FUNCTION auth.update_updated_at_column();
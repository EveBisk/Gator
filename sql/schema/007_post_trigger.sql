-- +goose Up
CREATE TRIGGER update_updated_at AFTER UPDATE
    ON posts
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at();

-- +goose Down
DROP TRIGGER IF EXISTS update_last_updated_at();
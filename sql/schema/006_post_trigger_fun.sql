-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS
$update_updated_at$
BEGIN
    NEW.updated_at := NOW();
    RETURN NEW;
END;  -- FUNCTION END
$update_updated_at$
LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
DROP FUNCTION IF EXISTS set_last_updated_at ON your_table;
-- DROP FUNCTION IF EXISTS update_last_updated_at();
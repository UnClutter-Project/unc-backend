-- +goose Up
-- +goose StatementBegin
ALTER TABLE clothing_category
    DROP COLUMN IF EXISTS user_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE clothing_category
    ADD COLUMN IF NOT EXISTS user_id UUID REFERENCES users(id) ON DELETE CASCADE;
-- +goose StatementEnd

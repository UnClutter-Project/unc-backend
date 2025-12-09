-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ALTER COLUMN is_verified SET NOT NULL;

ALTER TABLE users
    ALTER COLUMN is_verified SET DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    ALTER COLUMN is_verified DROP DEFAULT;

ALTER TABLE users
    ALTER COLUMN is_verified DROP NOT NULL;
-- +goose StatementEnd

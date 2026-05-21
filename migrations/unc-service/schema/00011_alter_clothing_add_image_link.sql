-- +goose Up
-- +goose StatementBegin
ALTER TABLE clothing
    ADD COLUMN image_link TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE clothing
    DROP COLUMN image_link;
-- +goose StatementEnd

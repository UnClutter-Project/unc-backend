-- +goose Up
-- +goose StatementBegin
CREATE TYPE gender_type AS ENUM ('male', 'female', 'prefer_not_to_say');

ALTER TABLE users
    ADD COLUMN gender gender_type,
    ADD COLUMN dob DATE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN gender,
    DROP COLUMN dob;

DROP TYPE gender_type;
-- +goose StatementEnd

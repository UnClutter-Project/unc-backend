-- +goose Up
-- +goose StatementBegin
CREATE TABLE color (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hex_value VARCHAR(7) NOT NULL,
    color_group_name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE color;
-- +goose StatementEnd

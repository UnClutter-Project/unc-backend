-- +goose Up
-- +goose StatementBegin
CREATE TABLE otp (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    token TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expired_at TIMESTAMPTZ NOT NULL DEFAULT NOW() + INTERVAL '5 minutes'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE otp;
-- +goose StatementEnd

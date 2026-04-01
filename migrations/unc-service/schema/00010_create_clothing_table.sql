-- +goose Up
-- +goose StatementBegin
CREATE TABLE clothing (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    main_color_1_id UUID REFERENCES color(id),
    main_color_2_id UUID REFERENCES color(id),
    accent_color_id UUID REFERENCES color(id),
    clothing_category_id UUID REFERENCES clothing_category(id),
    clothing_type_id UUID REFERENCES clothing_type(id),
    brand TEXT,
    style TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_clothing_user_main_color_1 ON clothing (user_id, main_color_1_id);
CREATE INDEX idx_clothing_user_main_color_2 ON clothing (user_id, main_color_2_id);
CREATE INDEX idx_clothing_user_accent_color ON clothing (user_id, accent_color_id);
CREATE INDEX idx_clothing_user_category ON clothing (user_id, clothing_category_id);
CREATE INDEX idx_clothing_user_type ON clothing (user_id, clothing_type_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE clothing;
-- +goose StatementEnd

-- name: GetColorByHex :one
SELECT * FROM colors
WHERE hex_value = $1
  AND deleted_at IS NULL
LIMIT 1;

-- name: CreateColor :one
INSERT INTO colors (
    id,
    hex_value,
    color_group_name,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    $1,
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: GetClothingCategoryByValueAndUserID :one
SELECT * FROM clothing_categories
WHERE user_id = $1
  AND value = $2
  AND deleted_at IS NULL
LIMIT 1;

-- name: CreateClothingCategory :one
INSERT INTO clothing_categories (
    id,
    user_id,
    value,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    $1,
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: GetClothingTypeByValueAndUserID :one
SELECT * FROM clothing_types
WHERE user_id = $1
  AND value = $2
  AND deleted_at IS NULL
LIMIT 1;

-- name: CreateClothingType :one
INSERT INTO clothing_types (
    id,
    user_id,
    value,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    $1,
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: CreateClothing :one
INSERT INTO clothings (
    id,
    user_id,
    main_color_1_id,
    main_color_2_id,
    accent_color_id,
    clothing_category_id,
    clothing_type_id,
    brand,
    style,
    image_link,
    created_at,
    updated_at
) VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    NOW(),
    NOW()
)
RETURNING *;
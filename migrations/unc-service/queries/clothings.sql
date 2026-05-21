-- name: GetColorByHex :one
SELECT * FROM color
WHERE hex_value = $1
  AND deleted_at IS NULL
LIMIT 1;

-- name: CreateColor :one
INSERT INTO color (
    hex_value,
    color_group_name
) VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetClothingCategoryByValueAndUserID :one
SELECT * FROM clothing_category
WHERE user_id = $1
  AND value = $2
  AND deleted_at IS NULL
LIMIT 1;

-- name: CreateClothingCategory :one
INSERT INTO clothing_category (
    user_id,
    value
) VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetClothingTypeByValueAndUserID :one
SELECT * FROM clothing_type
WHERE user_id = $1
  AND value = $2
  AND deleted_at IS NULL
LIMIT 1;

-- name: CreateClothingType :one
INSERT INTO clothing_type (
    user_id,
    value
) VALUES (
    $1,
    $2
)
RETURNING *;

-- name: CreateClothing :one
INSERT INTO clothing (
    user_id,
    main_color_1_id,
    main_color_2_id,
    accent_color_id,
    clothing_category_id,
    clothing_type_id,
    brand,
    style,
    image_link
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
)
RETURNING *;
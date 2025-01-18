-- name: CreateApp :one
INSERT INTO apps (created_by_user_id, name, created_at, updated_at)
VALUES ($1, $2, NOW(), NOW())
RETURNING *;

-- name: GetApp :one
SELECT * FROM apps WHERE id = $1;

-- name: ListUserApps :many
SELECT * FROM apps WHERE created_by_user_id = $1;

-- name: CreateAPIKey :one
INSERT INTO api_keys (app_id, key)
VALUES ($1, $2)
RETURNING *;

-- name: GetAPIKey :one

SELECT * FROM api_keys WHERE key = $1 AND is_active = true;

-- name: UpdateAPIKeyLastUsed :exec
UPDATE api_keys SET last_used = NOW() WHERE id = $1;

-- name: UpsertKeyValue :one
INSERT INTO key_values (app_id, key, value, updated_at)
VALUES ($1, $2, $3, NOW())
ON CONFLICT (app_id, key)
DO UPDATE SET value = EXCLUDED.value, updated_at = NOW()
RETURNING *;

-- name: GetKeyValue :one
SELECT * FROM key_values WHERE app_id = $1 AND key = $2;

-- name: ListKeyValues :many
SELECT * FROM key_values WHERE app_id = $1;

-- name: DeleteKeyValue :exec
DELETE FROM key_values WHERE app_id = $1 AND key = $2;


-- name: create :one
INSERT INTO languages
    (language_name, code)
VALUES
    ($1, $2)
RETURNING *;

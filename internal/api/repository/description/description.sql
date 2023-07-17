-- name: create :one
INSERT INTO description
    (word_id, description, level)
VALUES
    ($1, $2, $3)
RETURNING *;

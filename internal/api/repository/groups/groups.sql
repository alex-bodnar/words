-- name: create :one
INSERT INTO groups
    (name)
VALUES
    ($1)
RETURNING *;

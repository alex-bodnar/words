-- name: create :one
INSERT INTO words
    (word, transcription, language_id)
VALUES
    ($1, $2, $3)
RETURNING *;

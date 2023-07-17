-- name: create :one
INSERT INTO translations
    (translation, transcription, language_id)
VALUES
    ($1, $2, $3)
RETURNING *;

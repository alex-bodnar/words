-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS languages (
    id            SERIAL                   NOT NULL UNIQUE PRIMARY KEY,
    language_name TEXT                     NOT NULL UNIQUE,
    code          TEXT                     NOT NULL UNIQUE check (length(code) = 2),
    created_at    TIMESTAMP DEFAULT now()  NOT NULL,
    updated_at    TIMESTAMP DEFAULT now()  NOT NULL 
);

CREATE TABLE IF NOT EXISTS words (
    id            UUID     DEFAULT uuid_generate_v4() NOT NULL UNIQUE PRIMARY KEY,
    number        SERIAL                              NOT NULL,
    word          TEXT                                NOT NULL,
    transcription TEXT,
    language_id   INTEGER                             NOT NULL REFERENCES languages(id),
    word_ts       tsvector,
    created_at    TIMESTAMP DEFAULT now()             NOT NULL,
    updated_at    TIMESTAMP DEFAULT now()             NOT NULL,
    UNIQUE (word, language_id)
);

CREATE INDEX words_word_ts_idx ON words USING GIN (word_ts);

CREATE TABLE IF NOT EXISTS translations (
    id             UUID     DEFAULT uuid_generate_v4() NOT NULL UNIQUE PRIMARY KEY,
    number         SERIAL                              NOT NULL,
    translation    TEXT                                NOT NULL,
    transcription  TEXT,
    language_id    INTEGER                             NOT NULL REFERENCES languages(id),
    translation_ts tsvector,
    created_at     TIMESTAMP DEFAULT now()             NOT NULL,
    updated_at     TIMESTAMP DEFAULT now()             NOT NULL
);

CREATE INDEX translations_translation_ts_idx ON translations USING GIN (translation_ts);

CREATE TABLE IF NOT EXISTS word_translations (
    word_id        UUID                               NOT NULL REFERENCES words(id),
    translation_id UUID                               NOT NULL REFERENCES translations(id),
    UNIQUE (word_id, translation_id)
);

CREATE TYPE language_level AS ENUM ('A1', 'A2', 'B1', 'B2', 'C1', 'C2');

CREATE TABLE IF NOT EXISTS description (
    id            UUID      DEFAULT uuid_generate_v4() NOT NULL UNIQUE PRIMARY KEY,
    number        SERIAL                               NOT NULL,
    word_id       UUID                                 NOT NULL REFERENCES words(id),
    description   TEXT                                 NOT NULL,
    level         language_level                       NOT NULL,
    created_at    TIMESTAMP DEFAULT now()              NOT NULL,
    updated_at    TIMESTAMP DEFAULT now()              NOT NULL
);

CREATE TABLE IF NOT EXISTS groups (
    id         SERIAL                  NOT NULL UNIQUE PRIMARY KEY,
    name       TEXT                    NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT now() NOT NULL,
    updated_at TIMESTAMP DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS description_groups (
    description_id UUID    NOT NULL REFERENCES description(id),
    group_id       INTEGER NOT NULL REFERENCES groups(id),
    UNIQUE (description_id, group_id)
);

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION update_word_tsvector() RETURNS TRIGGER AS $$
BEGIN
    NEW.word_ts := to_tsvector((SELECT language_name FROM languages WHERE id = NEW.language_id)::regconfig, NEW.word);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_word_tsvector_trigger BEFORE INSERT OR UPDATE ON words
    FOR EACH ROW EXECUTE PROCEDURE update_word_tsvector();

CREATE TRIGGER update_translation_tsvector_trigger BEFORE INSERT OR UPDATE ON translations
    FOR EACH ROW EXECUTE PROCEDURE update_word_tsvector();


CREATE OR REPLACE FUNCTION set_updated_at() RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at := now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER languages_set_updated_at BEFORE UPDATE ON languages
    FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

CREATE TRIGGER words_set_updated_at BEFORE UPDATE ON words
    FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

CREATE TRIGGER translations_set_updated_at BEFORE UPDATE ON translations
    FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

CREATE TRIGGER description_set_updated_at BEFORE UPDATE ON description
    FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- +migrate StatementEnd

-- +migrate Down
DROP TRIGGER description_set_updated_at ON description;
DROP TRIGGER translations_set_updated_at ON translations;
DROP TRIGGER words_set_updated_at ON words;
DROP TRIGGER languages_set_updated_at ON languages;
DROP FUNCTION set_updated_at();

DROP TRIGGER update_translation_tsvector_trigger ON translations;
DROP TRIGGER update_word_tsvector_trigger ON words;
DROP FUNCTION update_word_tsvector();

DROP TABLE description_groups;
DROP TABLE groups;
DROP TABLE description;
DROP TYPE language_level;
DROP INDEX words_word_ts_idx;
DROP INDEX translations_translation_ts_idx;
DROP TABLE word_translations;
DROP TABLE translations;
DROP TABLE words;
DROP TABLE languages;

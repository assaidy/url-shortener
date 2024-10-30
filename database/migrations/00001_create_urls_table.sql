-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS urls (
    id INTEGER PRIMARY KEY AUTOINCREMENT, -- i don't use it. idk why tf i'm putting it here
    original_url TEXT NOT NULL UNIQUE,
    short_code TEXT NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT (DATETIME('now')),
    updated_at DATETIME NOT NULL DEFAULT (DATETIME('now')),
    access_count INTEGER DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE urls;
-- +goose StatementEnd


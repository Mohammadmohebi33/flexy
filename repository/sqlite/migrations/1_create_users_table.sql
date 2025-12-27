-- +migrate Up
CREATE TABLE users (
                       id INTEGER PRIMARY KEY AUTOINCREMENT,
                       name TEXT NOT NULL,
                       email TEXT UNIQUE,
                       email_verified_at DATETIME,
                       password TEXT NOT NULL,
                       score_week INTEGER NOT NULL DEFAULT 0,
                       score_month INTEGER NOT NULL DEFAULT 0,
                       score INTEGER NOT NULL DEFAULT 0,
                       avatar_url TEXT,
                       status TEXT NOT NULL DEFAULT 'active',
                       active_days INTEGER NOT NULL DEFAULT 0,
                       last_active_at DATETIME,
                       remember_token TEXT,
                       is_verified INTEGER NOT NULL DEFAULT 0,
                       is_admin INTEGER NOT NULL DEFAULT 0,
                       is_banned INTEGER NOT NULL DEFAULT 0,
                       fcm_token TEXT,
                       created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                       updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_email ON users(email);
CREATE INDEX idx_status ON users(status);

-- +migrate Down
DROP TABLE IF EXISTS users;
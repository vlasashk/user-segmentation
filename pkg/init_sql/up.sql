CREATE TABLE IF NOT EXISTS users
(
    "id"  BIGSERIAL PRIMARY KEY,
    "user_id" BIGINT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS segments
(
    "id"   BIGSERIAL PRIMARY KEY,
    "slug" varchar(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS user_segments
(
    user_id    BIGINT REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    segment_id BIGINT REFERENCES segments (id) ON DELETE CASCADE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    PRIMARY KEY (user_id, segment_id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_user_segments_not_deleted
    ON user_segments (user_id, segment_id)
    WHERE deleted_at IS NULL;
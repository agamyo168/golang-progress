CREATE TABLE IF NOT EXISTS posts(
    id bigserial PRIMARY KEY,
    content varchar(255) NOT NULL,
    user_id bigint NOT NULL,
    tags TEXT[] DEFAULT '{}',
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_posts_modtime
    BEFORE UPDATE ON posts
    FOR EACH ROW
    EXECUTE PROCEDURE update_modified_column();
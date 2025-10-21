CREATE TABLE IF NOT EXISTS profiles (
                                        id SERIAL PRIMARY KEY,
                                        user_id INT UNIQUE NOT NULL REFERENCES users(id),
    display_name TEXT,
    bio TEXT,
    tags TEXT[],
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
    );

CREATE INDEX IF NOT EXISTS idx_profiles_tags ON profiles USING GIN (tags);

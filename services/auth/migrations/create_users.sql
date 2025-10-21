CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     email TEXT UNIQUE NOT NULL,
                                     password_hash TEXT NOT NULL,
                                     role TEXT NOT NULL CHECK (role IN ('creator', 'advertiser', 'admin')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
    );

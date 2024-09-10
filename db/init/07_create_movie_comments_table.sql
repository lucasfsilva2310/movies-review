CREATE TABLE IF NOT EXISTS movie_comments (
    id SERIAL PRIMARY KEY,
    movie_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    comment VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    --Foreign Keys
    FOREIGN KEY (movie_id) REFERENCES movies(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
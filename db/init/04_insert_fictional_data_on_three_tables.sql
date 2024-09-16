-- Inserting data on movies
INSERT INTO movies (title, description, release_date, tags, platforms)
VALUES
('Inception', 'A mind-bending thriller', '2010-07-16', '["Sci-Fi", "Action"]', '["Netflix", "Amazon"]'),
('The Matrix', 'A hacker discovers a hidden reality', '1999-03-31', '["Sci-Fi", "Action"]', '["Hulu", "Netflix"]'),
('The Godfather', 'An epic story of a crime family', '1972-03-24', '["Drama", "Crime"]', '["Amazon", "Apple"]'),
('Forrest Gump', 'A man recounts his extraordinary life', '1994-07-06', '["Drama", "Romance"]', '["Netflix", "Disney+"]'),
('The Dark Knight', 'Batman faces the Joker', '2008-07-18', '["Action", "Drama"]', '["Netflix", "Hulu"]');

-- Inserting data on users
INSERT INTO users (username, password, full_name, email)
VALUES
('john_doe', '$2a$10$47dAe7Vodpmou/XWjiJC0.xrYrOMfl7ck10lajLkbkdgJsp898lwW', 'John Doe', 'john@example.com'),
('jane_smith', '$2a$10$47dAe7Vodpmou/XWjiJC0.xrYrOMfl7ck10lajLkbkdgJsp898lwW', 'Jane Smith', 'jane@example.com'),
('alice_wonder', '$2a$10$47dAe7Vodpmou/XWjiJC0.xrYrOMfl7ck10lajLkbkdgJsp898lwW', 'Alice Wonder', 'alice@example.com'),
('bob_builder', '$2a$10$47dAe7Vodpmou/XWjiJC0.xrYrOMfl7ck10lajLkbkdgJsp898lwW', 'Bob Builder', 'bob@example.com'),
('charlie_brown', 'password123', 'Charlie Brown', 'charlie@example.com');

-- Inserting data on ratings
INSERT INTO ratings (movie_id, user_id, score)
VALUES
(1, 1, 5), 
(2, 2, 4), 
(3, 3, 5), 
(4, 4, 3), 
(5, 5, 5), 
(1, 2, 4), 
(2, 3, 3), 
(3, 4, 4), 
(4, 5, 5), 
(5, 1, 4); 
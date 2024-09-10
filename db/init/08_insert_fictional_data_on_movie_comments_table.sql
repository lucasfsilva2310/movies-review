-- Inserting on movie_comments table

-- user 1 comment movies 1 e 2
INSERT INTO movie_comments (movie_id, user_id, comment, created_at, updated_at)
VALUES 
    (1, 1, 'Amazing movie, loved the action!', NOW(), NOW()),
    (2, 1, 'Great storyline, but a bit slow.', NOW(), NOW());

-- user 2 comment movie 3
INSERT INTO movie_comments (movie_id, user_id, comment, created_at, updated_at)
VALUES 
    (3, 2, 'One of the best movies I have seen!', NOW(), NOW());

-- user 3 comment movies 1 e 5
INSERT INTO movie_comments (movie_id, user_id, comment, created_at, updated_at)
VALUES 
    (1, 3, 'The action sequences were outstanding.', NOW(), NOW()),
    (5, 3, 'Not bad, but expected more from the ending.', NOW(), NOW());

-- user 4 comment movie 2
INSERT INTO movie_comments (movie_id, user_id, comment, created_at, updated_at)
VALUES 
    (2, 4, 'Really loved the characters and the plot twists!', NOW(), NOW());

-- user 5 comment movies 3 e 4
INSERT INTO movie_comments (movie_id, user_id, comment, created_at, updated_at)
VALUES 
    (3, 5, 'A masterpiece, highly recommend!', NOW(), NOW()),
    (4, 5, 'Good movie, but too long for my taste.', NOW(), NOW());

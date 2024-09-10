
--Inserting data on watched_movies

-- user 1 watched movies 1, 2 e 3
INSERT INTO watched_movies (movie_id, user_id, watched, created_at, updated_at)
VALUES 
    (1, 1, true, NOW(), NOW()),  
    (2, 1, true, NOW(), NOW()),  
    (3, 1, true, NOW(), NOW()); 

--  user 2 watched movie 1 e 4
INSERT INTO watched_movies (movie_id, user_id, watched, created_at, updated_at)
VALUES 
    (1, 2, true, NOW(), NOW()),
    (4, 2, true, NOW(), NOW());

--  user 3 watched movie 5 then marked as unwatched
INSERT INTO watched_movies (movie_id, user_id, watched, created_at, updated_at)
VALUES 
    (5, 3, false, NOW(), NOW());

--  user 4 watched movie 2 e 3
INSERT INTO watched_movies (movie_id, user_id, watched, created_at, updated_at)
VALUES 
    (2, 4, true, NOW(), NOW()),
    (3, 4, true, NOW(), NOW());

--  user 5 watched movie 1 then marked movie 2 as unwatched
INSERT INTO watched_movies (movie_id, user_id, watched, created_at, updated_at)
VALUES 
    (1, 5, true, NOW(), NOW()),
    (2, 5, false, NOW(), NOW());

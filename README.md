# Movies Review

This is an API that will have a list of movies.

Each movie can be rated one time and can have a single comment for every user.

A user can also list any number of movies as `watched` to have a tracking system of every watched movie.

Right now you can comment and rate a movie without watching it.

## Features

### User

- A new user can be registered by providing a username, password, full name and an email.
- To access every other endpoint besides `/login` and `/register`, the user needs to have an access token granted by requesting a login successfully.

### Movies

- A movie can be created by any user.
- A user cannot create a new movie with an existing movie name already created.
- A user can delete a movie. It does not have a validation to check if the same user created the deleted movie.
- A user can list all movies or get a specific movie by ID

### Ratings

- A movie can be rated one time by every user.
- The rate can have a integer range between 1 to 5.
- A rating can be deleted by a user if user has same id as user_id column

### Watched Movies

- A movie can be listed as watched one time by every user.
- A user can change the `watched` flag to true or false whenever he or she feels like it
- A watched movie can be deleted by a user if user has same id as user_id column.

### Movie Comments

- A movie can have a single comment for every user.
- It does not have and endpoint to patch the current comment row for now.
- A movie comment can be deleted by a user if user has same id as user_id column.

# How to run API

To run the full API, including a local database being runned by docker, you just need one line of command:

```
docker-compose up

or

docker compose up
```

You will need docker compose installed locally of course.

It will automatically build both API and database containers locally and run, already connecting the API to the database and auto refreshing the application with every code change by using air.

# API Documentation

<details>
<summary>Documentation by Insomnia</summary>

```
[
  {
    "url": "http://localhost:8000/hello",
    "name": "HELLO WORLD",
    "method": "GET",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/watched-movies/user/:id",
    "name": "GET ALL WATCHED MOVIES BY USER ID",
    "method": "GET",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/watched-movies/movie/:id",
    "name": "GET ALL WATCHED MOVIES BY MOVIE ID",
    "method": "GET",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/watched-movies/:id",
    "name": "DELETE WATCHED MOVIE",
    "method": "DELETE",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/watched-movies",
    "name": "GET ALL WATCHED MOVIES",
    "method": "GET",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/watched-movies",
    "name": "CREATE WATCHED MOVIE",
    "method": "POST",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/movie-comments/user/:id",
    "name": "GET ALL MOVIE COMMENTS BY USER ID",
    "method": "GET",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/movie-comments/:id",
    "name": "DELETE MOVIE COMMENTS",
    "method": "DELETE",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/movie-comments/movie/:id",
    "name": "GET ALL MOVIE COMMENTS BY MOVIE ID",
    "method": "GET",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/movie-comments",
    "name": "GET ALL MOVIE COMMENTS",
    "method": "GET",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/movie-comments",
    "name": "CREATE MOVIE COMMENTS",
    "method": "POST",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/ratings",
    "name": "CREATE RATING",
    "method": "POST",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/ratings",
    "name": "GET ALL RATINGS",
    "method": "GET",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/ratings/movie/:id",
    "name": "GET ALL RATINGS BY MOVIE ID",
    "method": "GET",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/ratings/:id",
    "name": "DELETE RATINGS",
    "method": "DELETE",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/ratings/user/:id",
    "name": "GET ALL RATINGS BY USER ID",
    "method": "GET",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/users",
    "name": "CREATE USER",
    "method": "POST",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/users/:id",
    "name": "GET USER BY ID",
    "method": "GET",
    "authentication": {
      "type": "apikey",
      "disabled": false,
      "key": "Authorization",
      "value": "Bearer yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjYxODE5NDQsInVzZXJuYW1lIjoibHVjYXNmc2lsdmEifQ.8lCg3Z0IwvVBnHivg9qRM2d81voUYSD8o9Zg689jFPs",
      "addTo": "header"
    }
  },
  {
    "url": "http://localhost:8000/users/:id",
    "name": "DELETE USER ",
    "method": "DELETE",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/users/:username",
    "name": "DELETE USER BY USERNAME",
    "method": "DELETE",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/users/username/:username",
    "name": "GET USER BY USERNAME",
    "method": "GET",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/auth/register",
    "name": "REGISTER",
    "method": "POST",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/auth/login",
    "name": "LOGIN",
    "method": "POST",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/movies/:id",
    "name": "GET MOVIE BY ID",
    "method": "GET",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/movies",
    "name": "CREATE MOVIE",
    "method": "POST",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/movies/:id",
    "name": "DELETE MOVIE",
    "method": "DELETE",
    "authentication": "User Authentication"
  },
  {
    "url": "http://localhost:8000/movies",
    "name": "GET ALL MOVIES",
    "method": "GET",
    "authentication": "User Authentication"
  },

]
```

</details>

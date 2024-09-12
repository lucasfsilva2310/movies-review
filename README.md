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
- A user cannot create a new movie with an existing movie already created.
- A user can delete a movie. It does not have a validation to check if the same user created the deleted movie.
- A user can list all movies or get a specific movie by ID

### Ratings

- A movie can be rated one time by every user.
- The rate can have a integer range between 1 to 5.
- Every rating can be deleted by any user.

### Watched Movies

- A movie can be listed as watched one time by every user.
- It does not have an endpoint to patch the current watched movie row as `watched = false` for now.
- A watched movie can be deleted by every user.

### Movie Comments

- A movie can have a single comment for every user.
- It does not have and endpoint to patch the current comment row for now.
- A movie comment can be deleted by every user.

# How to run API

To run the full API, including a local database being runned by docker, you just need one line of command:

```
docker-compose up

or

docker compose up
```

It will automatically build both API and database containers locally and run, already connecing the API to the database.

# API Documentation

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

<details>
<summary>Documentation by Insomnia</summary>

```
{
  "resources": [
    {
      "_id": "req_de0e031898f14558990da6a4125feb91",
      "parentId": "wrk_d80a3899a3864e9789f276b0e0fbbc4f",
      "modified": 1726011400419,
      "created": 1723034542887,
      "url": "http://localhost:8000/hello",
      "name": "HELLO WORLD",
      "description": "",
      "method": "GET",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1726010191732,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "wrk_d80a3899a3864e9789f276b0e0fbbc4f",
      "parentId": null,
      "modified": 1723496590267,
      "created": 1723034539144,
      "name": "Movies-review-API-Go-Gin",
      "description": "",
      "scope": "collection",
      "_type": "workspace"
    },
    {
      "_id": "req_404091f63cb64386b389960ba15efda2",
      "parentId": "fld_71f7ae4bf55f457384b1d8dcf9b5f271",
      "modified": 1726095623887,
      "created": 1726010225432,
      "url": "http://localhost:8000/watched-movies/user/1",
      "name": "GET ALL WATCHED MOVIES BY USER ID",
      "description": "",
      "method": "GET",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {
        "type": "apikey",
        "disabled": false,
        "key": "",
        "value": "",
        "addTo": "header"
      },
      "metaSortKey": -1726010204718,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "fld_71f7ae4bf55f457384b1d8dcf9b5f271",
      "parentId": "wrk_d80a3899a3864e9789f276b0e0fbbc4f",
      "modified": 1726011335239,
      "created": 1726010191632,
      "name": "watched_movies",
      "description": "",
      "environment": {},
      "environmentPropertyOrder": null,
      "metaSortKey": -1726010191632,
      "_type": "request_group"
    },
    {
      "_id": "req_7a28a687e88b41e1953fb8d3c30f07e2",
      "parentId": "fld_71f7ae4bf55f457384b1d8dcf9b5f271",
      "modified": 1726010248235,
      "created": 1726010220076,
      "url": "http://localhost:8000/watched-movies/movie/1",
      "name": "GET ALL WATCHED MOVIES BY MOVIE ID",
      "description": "",
      "method": "GET",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1726010204618,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_8b0cb288cd9a4aaa836af2a33630791a",
      "parentId": "fld_71f7ae4bf55f457384b1d8dcf9b5f271",
      "modified": 1726179919649,
      "created": 1726179912595,
      "url": "http://localhost:8000/watched-movies/1",
      "name": "DELETE WATCHED MOVIE",
      "description": "",
      "method": "DELETE",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1726010204568,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_097945c97be9436d85cbe57052c431b6",
      "parentId": "fld_71f7ae4bf55f457384b1d8dcf9b5f271",
      "modified": 1726010250001,
      "created": 1726010211067,
      "url": "http://localhost:8000/watched-movies",
      "name": "GET ALL WATCHED MOVIES",
      "description": "",
      "method": "GET",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1726010204518,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_ffc5c0a2dbe74dff88cabb944313d546",
      "parentId": "fld_71f7ae4bf55f457384b1d8dcf9b5f271",
      "modified": 1726010251681,
      "created": 1726010201662,
      "url": "http://localhost:8000/watched-movies",
      "name": "CREATE WATCHED MOVIE",
      "description": "",
      "method": "POST",
      "body": {
        "mimeType": "application/json",
        "text": "{\n\t\"score\": 2,\n\t\"movie_id\": 1,\n\t\"user_id\": 1\n}"
      },
      "parameters": [],
      "headers": [
        {
          "name": "Content-Type",
          "value": "application/json"
        },
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1726010204418,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_e3a18cbcd49c43cb9da75f9b26030611",
      "parentId": "fld_48ecd61b488046b78477253a9a4ad47f",
      "modified": 1726011406598,
      "created": 1726011323172,
      "url": "http://localhost:8000/movie-comments/user/1",
      "name": "GET ALL MOVIE COMMENTS BY USER ID",
      "description": "",
      "method": "GET",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1726010204718,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "fld_48ecd61b488046b78477253a9a4ad47f",
      "parentId": "wrk_d80a3899a3864e9789f276b0e0fbbc4f",
      "modified": 1726011323165,
      "created": 1726011323165,
      "name": "movie_comments",
      "description": "",
      "environment": {},
      "environmentPropertyOrder": null,
      "metaSortKey": -1725793499561.5,
      "_type": "request_group"
    },
    {
      "_id": "req_a7f2eef00c3d4191a9706098aae663b5",
      "parentId": "fld_48ecd61b488046b78477253a9a4ad47f",
      "modified": 1726179845658,
      "created": 1726179838554,
      "url": "http://localhost:8000/movie-comments/1",
      "name": "DELETE MOVIE COMMENTS",
      "description": "",
      "method": "DELETE",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1726010204668,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_2bdca72086ac4bc581a2a5892ef83d4e",
      "parentId": "fld_48ecd61b488046b78477253a9a4ad47f",
      "modified": 1726011411156,
      "created": 1726011323170,
      "url": "http://localhost:8000/movie-comments/movie/1",
      "name": "GET ALL MOVIE COMMENTS BY MOVIE ID",
      "description": "",
      "method": "GET",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1726010204618,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_f722864e8fe74a93817fa58313837745",
      "parentId": "fld_48ecd61b488046b78477253a9a4ad47f",
      "modified": 1726011413641,
      "created": 1726011323169,
      "url": "http://localhost:8000/movie-comments",
      "name": "GET ALL MOVIE COMMENTS",
      "description": "",
      "method": "GET",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1726010204518,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_32e6c967373c4379a58d0fe2c678cebf",
      "parentId": "fld_48ecd61b488046b78477253a9a4ad47f",
      "modified": 1726011415975,
      "created": 1726011323167,
      "url": "http://localhost:8000/movie-comments",
      "name": "CREATE MOVIE COMMENTS",
      "description": "",
      "method": "POST",
      "body": {
        "mimeType": "application/json",
        "text": "{\n\t\"score\": 2,\n\t\"movie_id\": 1,\n\t\"user_id\": 1\n}"
      },
      "parameters": [],
      "headers": [
        {
          "name": "Content-Type",
          "value": "application/json"
        },
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1726010204418,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_64a23555d1314eb4861e799c058bea12",
      "parentId": "fld_dc91f085aef343dca73b0a648956fb24",
      "modified": 1725577604190,
      "created": 1725576810777,
      "url": "http://localhost:8000/ratings",
      "name": "CREATE RATING",
      "description": "",
      "method": "POST",
      "body": {
        "mimeType": "application/json",
        "text": "{\n\t\"score\": 2,\n\t\"movie_id\": 1,\n\t\"user_id\": 1\n}"
      },
      "parameters": [],
      "headers": [
        {
          "name": "Content-Type",
          "value": "application/json"
        },
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1725576810777,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "fld_dc91f085aef343dca73b0a648956fb24",
      "parentId": "wrk_d80a3899a3864e9789f276b0e0fbbc4f",
      "modified": 1726011384312,
      "created": 1725576807491,
      "name": "ratings",
      "description": "",
      "environment": {},
      "environmentPropertyOrder": null,
      "metaSortKey": -1725576807491,
      "_type": "request_group"
    },
    {
      "_id": "req_e599c1d843174d9ab2a763c4bb962f06",
      "parentId": "fld_dc91f085aef343dca73b0a648956fb24",
      "modified": 1725577570229,
      "created": 1725576873666,
      "url": "http://localhost:8000/ratings",
      "name": "GET ALL RATINGS",
      "description": "",
      "method": "GET",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1725556796924.5,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_7cd412af691c417f9a257e6791a46770",
      "parentId": "fld_dc91f085aef343dca73b0a648956fb24",
      "modified": 1726008557347,
      "created": 1726008554013,
      "url": "http://localhost:8000/ratings/movie/1",
      "name": "GET ALL RATINGS BY MOVIE ID",
      "description": "",
      "method": "GET",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1725546789998.25,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_b951ca8b6e874dec8918c1f4eb5b0b93",
      "parentId": "fld_dc91f085aef343dca73b0a648956fb24",
      "modified": 1726179975913,
      "created": 1726179971009,
      "url": "http://localhost:8000/ratings/1",
      "name": "DELETE RATINGS",
      "description": "",
      "method": "DELETE",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1725544288266.6875,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_854a234648db42308fc4f0f64047e5fd",
      "parentId": "fld_dc91f085aef343dca73b0a648956fb24",
      "modified": 1726008572443,
      "created": 1726008570601,
      "url": "http://localhost:8000/ratings/user/1",
      "name": "GET ALL RATINGS BY USER ID",
      "description": "",
      "method": "GET",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1725541786535.125,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_495147ed9fbf4a889e1908358ff5f5e8",
      "parentId": "fld_a66bbd36968943fd9a4f1203d26be3a2",
      "modified": 1725536899956,
      "created": 1725536783072,
      "url": "http://localhost:8000/users",
      "name": "CREATE USER",
      "description": "",
      "method": "POST",
      "body": {
        "mimeType": "application/json",
        "text": "{\n\t\"username\": \"lucasfsilva\",\n\t\"password\":\"password\",\n\t\"full_name\": \"Lucas Ferreira Silva\",\n\t\"email\": \"lucasfsilva2310@gmail.com\"\n}"
      },
      "parameters": [],
      "headers": [
        {
          "name": "Content-Type",
          "value": "application/json"
        },
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1725536783072,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "fld_a66bbd36968943fd9a4f1203d26be3a2",
      "parentId": "wrk_d80a3899a3864e9789f276b0e0fbbc4f",
      "modified": 1726011392770,
      "created": 1725536779419,
      "name": "users",
      "description": "",
      "environment": {},
      "environmentPropertyOrder": null,
      "metaSortKey": -1725536779419,
      "_type": "request_group"
    },
    {
      "_id": "req_5c9b74c2127644f4adeb8307eea9ae6a",
      "parentId": "fld_a66bbd36968943fd9a4f1203d26be3a2",
      "modified": 1726095587304,
      "created": 1725539436023,
      "url": "http://localhost:8000/users/2",
      "name": "GET USER BY ID",
      "description": "",
      "method": "GET",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {
        "type": "apikey",
        "disabled": false,
        "key": "Authorization",
        "value": "Bearer yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjYxODE5NDQsInVzZXJuYW1lIjoibHVjYXNmc2lsdmEifQ.8lCg3Z0IwvVBnHivg9qRM2d81voUYSD8o9Zg689jFPs",
        "addTo": "header"
      },
      "metaSortKey": -1725536781295.5,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_9c2d962854584a7592510b89cab888c6",
      "parentId": "fld_a66bbd36968943fd9a4f1203d26be3a2",
      "modified": 1726180005877,
      "created": 1726180002957,
      "url": "http://localhost:8000/users/2",
      "name": "DELETE USER ",
      "description": "",
      "method": "DELETE",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {
        "type": "apikey",
        "disabled": false,
        "key": "Authorization",
        "value": "Bearer yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjYxODE5NDQsInVzZXJuYW1lIjoibHVjYXNmc2lsdmEifQ.8lCg3Z0IwvVBnHivg9qRM2d81voUYSD8o9Zg689jFPs",
        "addTo": "header"
      },
      "metaSortKey": -1725536780851.375,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_0a641995a5f24e8183f7c2b49b516115",
      "parentId": "fld_a66bbd36968943fd9a4f1203d26be3a2",
      "modified": 1726180026107,
      "created": 1726180016316,
      "url": "http://localhost:8000/users/john_doe",
      "name": "DELETE USER BY USERNAME",
      "description": "",
      "method": "DELETE",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {
        "type": "apikey",
        "disabled": false,
        "key": "Authorization",
        "value": "Bearer yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjYxODE5NDQsInVzZXJuYW1lIjoibHVjYXNmc2lsdmEifQ.8lCg3Z0IwvVBnHivg9qRM2d81voUYSD8o9Zg689jFPs",
        "addTo": "header"
      },
      "metaSortKey": -1725536780629.3125,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_0f83b0ef19eb41899970cde03b0ac4a4",
      "parentId": "fld_a66bbd36968943fd9a4f1203d26be3a2",
      "modified": 1726008521013,
      "created": 1725539874983,
      "url": "http://localhost:8000/users/username/john_doe",
      "name": "GET USER BY USERNAME",
      "description": "",
      "method": "GET",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1725536780407.25,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_faa09f3efd034e06a3eae0c1af029b2b",
      "parentId": "fld_96d17ca29da94c09b6e42d327670ddd6",
      "modified": 1726094697179,
      "created": 1726094662691,
      "url": "http://localhost:8000/auth/register",
      "name": "REGISTER",
      "description": "",
      "method": "POST",
      "body": {
        "mimeType": "application/json",
        "text": "{\n\t\"username\": \"lucasfsilva\",\n\t\"password\":\"password\",\n\t\"full_name\": \"Lucas Ferreira Silva\",\n\t\"email\": \"lucasfsilva2310@gmail.com\"\n}"
      },
      "parameters": [],
      "headers": [
        {
          "name": "Content-Type",
          "value": "application/json"
        },
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1725536783072,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "fld_96d17ca29da94c09b6e42d327670ddd6",
      "parentId": "wrk_d80a3899a3864e9789f276b0e0fbbc4f",
      "modified": 1726094662689,
      "created": 1726094662689,
      "name": "auth",
      "description": "",
      "environment": {},
      "environmentPropertyOrder": null,
      "metaSortKey": -1725536769786,
      "_type": "request_group"
    },
    {
      "_id": "req_7f64401ac06947d0a7be2f0f9f516b1e",
      "parentId": "fld_96d17ca29da94c09b6e42d327670ddd6",
      "modified": 1726094944261,
      "created": 1726094688451,
      "url": "http://localhost:8000/auth/login",
      "name": "LOGIN",
      "description": "",
      "method": "POST",
      "body": {
        "mimeType": "application/json",
        "text": "{\n\t\"username\": \"lucasfsilva\",\n\t\"password\":\"password\"\n}"
      },
      "parameters": [],
      "headers": [
        {
          "name": "Content-Type",
          "value": "application/json"
        },
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1725536782183.75,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_c7222504bf9348a7a84b274c08dd90c4",
      "parentId": "fld_8ad672cb945d42e9b8c11e57c9474550",
      "modified": 1725536768656,
      "created": 1724799791330,
      "url": "http://localhost:8000/movies/123",
      "name": "GET MOVIE BY ID",
      "description": "",
      "method": "GET",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1725536764238,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "fld_8ad672cb945d42e9b8c11e57c9474550",
      "parentId": "wrk_d80a3899a3864e9789f276b0e0fbbc4f",
      "modified": 1726011396044,
      "created": 1725536760152,
      "name": "movies",
      "description": "",
      "environment": {},
      "environmentPropertyOrder": null,
      "metaSortKey": -1725536760153,
      "_type": "request_group"
    },
    {
      "_id": "req_1a1e8c19bfa4484881e09c1ada7f1689",
      "parentId": "fld_8ad672cb945d42e9b8c11e57c9474550",
      "modified": 1725538206573,
      "created": 1724860439920,
      "url": "http://localhost:8000/movies",
      "name": "CREATE MOVIE",
      "description": "",
      "method": "POST",
      "body": {
        "mimeType": "application/json",
        "text": "{\n\t\"title\": \"Lord Of The Rings\",\n\t\"description\": \"Story about a ring and a lord\",\n\t\"releaseDate\": \"1993/02/28\",\n\t\"tags\": [\"Fantasy\"],\n\t\"platforms\": [\"Netflix\"],\n\t\"createdAt\": \"{% faker 'isoTimestamp' %}\",\n\t\"updatedAt\": \"{% faker 'isoTimestamp' %}\"\n}"
      },
      "parameters": [],
      "headers": [
        {
          "name": "Content-Type",
          "value": "application/json"
        },
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1725536764213,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_3c85895806794a5c9e15b8059aba37a2",
      "parentId": "fld_8ad672cb945d42e9b8c11e57c9474550",
      "modified": 1725536773198,
      "created": 1724947946201,
      "url": "http://localhost:8000/movies/7",
      "name": "DELETE MOVIE",
      "description": "",
      "method": "DELETE",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1725536764200.5,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "req_4a078e5fe2b14485afb1cc9cf0620874",
      "parentId": "fld_8ad672cb945d42e9b8c11e57c9474550",
      "modified": 1725536767170,
      "created": 1724799776645,
      "url": "http://localhost:8000/movies",
      "name": "GET ALL MOVIES",
      "description": "",
      "method": "GET",
      "body": {},
      "parameters": [],
      "headers": [
        {
          "name": "User-Agent",
          "value": "insomnia/9.3.3"
        }
      ],
      "authentication": {},
      "metaSortKey": -1725536764188,
      "isPrivate": false,
      "pathParameters": [],
      "settingStoreCookies": true,
      "settingSendCookies": true,
      "settingDisableRenderRequestBody": false,
      "settingEncodeUrl": true,
      "settingRebuildPath": true,
      "settingFollowRedirects": "global",
      "_type": "request"
    },
    {
      "_id": "env_4962c87baf257a831f4aa6b75510ed87e93232cb",
      "parentId": "wrk_d80a3899a3864e9789f276b0e0fbbc4f",
      "modified": 1723034539147,
      "created": 1723034539147,
      "name": "Base Environment",
      "data": {},
      "dataPropertyOrder": null,
      "color": null,
      "isPrivate": false,
      "metaSortKey": 1723034539147,
      "_type": "environment"
    },
    {
      "_id": "jar_4962c87baf257a831f4aa6b75510ed87e93232cb",
      "parentId": "wrk_d80a3899a3864e9789f276b0e0fbbc4f",
      "modified": 1723034539148,
      "created": 1723034539148,
      "name": "Default Jar",
      "cookies": [],
      "_type": "cookie_jar"
    }
  ]
}
```

</details>

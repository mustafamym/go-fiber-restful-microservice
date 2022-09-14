![fiber_cover_gh](https://user-images.githubusercontent.com/11155743/112001218-cf258b00-8b2f-11eb-9c6d-d6c38a09af86.jpg)

# go-fiber-restful-microservice
Build a RESTful API on 

Go

Fiber

PostgreSQL

JWT 

Swagger

Docker 

## Public API 
GET: /api/books, get all books;

GET: /api/book/{id}, get book by given ID;

GET: /api/token/new, create a new access token (for a demo);

## secured API (JWT protected):

POST: /api/secured/book, create a new book;

PATCH: /api/secured/book, update an existing book;

DELETE: /api/secured/book, delete an existing book;

##swagger

 Go to your API Docs page: [127.0.0.1:8080/swagger/index.html](http://127.0.0.1:8080/swagger/index.html)

![swaggerUI](/swaggerUI.png)



## Quick start

1. Rename `.env.example` to `.env` and fill it with your environment values.
2. Install [Docker](https://www.docker.com/get-started) and [migrate](https://github.com/golang-migrate/migrate) tool for applying migrations.
3. Run project by this command:

```bash
make docker.run

# Process:
#   - Generate API docs by Swagger
#   - Create a new Docker network for containers
#   - Build and run Docker containers (Fiber, PostgreSQL)
#   - Apply database migrations (using github.com/golang-migrate/migrate)
```

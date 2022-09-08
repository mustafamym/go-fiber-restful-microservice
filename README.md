# go-fiber-restful-microservice
Build a RESTful API on Go: Fiber, PostgreSQL, JWT and Swagger docs in isolated Docker containers

## API methods
Public:

GET: /api/v1/books, get all books;
GET: /api/v1/book/{id}, get book by given ID;
GET: /api/v1/token/new, create a new access token (for a demo);

## Private (JWT protected):

POST: /api/v1/book, create a new book;
PATCH: /api/v1/book, update an existing book;
DELETE: /api/v1/book, delete an existing book;

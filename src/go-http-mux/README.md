# Simple GO Lang REST API

Simple RESTful API to create, read, update and delete books. No database implementation yet

## Quick Start
```
# Install mux router
go get -u github.com/gorilla/mux

go build
./http-api
```
## A powerful HTTP router and URL matcher for building Go web servers

https://github.com/gorilla/mux

## Endpoints

- GET api/books
- GET api/books/{id}
- DELETE api/books/{id}
- POST api/books
```
# Request sample
# {
#   "isbn":"4545454",
#   "title":"Book Three",
#   "author":{"firstname":"Harry",  "lastname":"White"}
# }
```
- PUT api/books/{id}

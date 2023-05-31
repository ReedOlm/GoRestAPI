# GoRestAPI
A simple REST API in Go using the Gin framework to manage Users with an in-memory database

## Instructions for Running:
  - clone repo
  - cd into repo and run "go run main.go"
  - in tool of choice (I've provided sample curl commands in testCommands.txt) hit localhost:8080 with any of the following endpoints

### Endpoints
  - GET     /users
  - GET     /users/:id
  - GET     /users/:id/friends
  - POST    /users
  - POST    /users/:id/friends
  - PUT     /users/:id
  - DELETE  /users/:id
  - DELETE  /users/:id/friends/:friendId

### Data format
  - :id and :friendId: any number convertable to int64
  - POST/PUT, json with this format:
```
{
    ID        int64 
    Name      string
    Email     string
    Age       int
    Friends   []int64
}
```
## Assumptions:
### General:
  - Use of framework allowed, not building server from scratch.
  - User ID's are unique and immutable.
  - For simplicity's sake, because no acceptance system exists, users are allowed to be friends with someone not following them, essentially creating a follower system.
### GET:
  - An empty database, or having no friends is a successful return.
### POST:
  - No validation on email / name validity required.
  - Friend adding request body formatted as a User, not validating whether new friend ID exists.
  - Adding a friend to list, will not add you to their list.
### PUT:
  - Requests will contain a complete User in body of request, all fields will be overwritten by request body object.
### DELETE:
  - Deleting a friend from list, will not delete you from their list.

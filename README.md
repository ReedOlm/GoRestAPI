# GoRestAPI
A simple REST API in Go to manage Users

Assumptions:
General:
  Use of framework allowed, not building server from scratch.
  User ID's are unique and immutable.
  For simplicity's sake, because no acceptance system exists, users are all friends with someone, who is in turn not friends with them, essentially creating a follower system.
GET:
  An empty database, or having no friends is a successful return.
POST:
  No validation on email / name validity required.
  Adding a friend to list, will not add you to their list.
PUT:
  Requests will contain a complete User in body of request, all fields will be overwritten by request body object.
DELETE:
  Deleting a friend from list, will not delete you from their list.

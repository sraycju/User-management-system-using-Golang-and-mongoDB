# User Management System Using Golang and mongoDB

Performs CRUD operations on MongoDB written in Golang. You can create, read, update, and delete users from the MongoDB instance using http requests.

### Features!

  - Performs CRUD operations on MongoDB
  - Used mux for routing
  - Follows model, dao, controller, router structure

### Project structure
```bash

├── README.md
│   ├── controller                  // API controllers / handlers
│   │   ├── userController.go
│   ├── models                    // DB models
│   │   └── userModel.go
├── router
│   └── userRouter.go                 // API routes
├── go.mod                      // Go modules
├── go.sum
├── main.go

```
### Installation

golang-mongodb-restful-starter-kit requires [Go](https://golang.org/) 1.10+ to run.

Install the dependencies and devDependencies and start the server.

```sh
$ git clone https://github.com/sraycju/User-management-system-using-Golang-and-mongoDB.git
$ cd  User-management-system-using-Golang-and-mongoDB
$ go run main.go
```
#### Building for source
For production release:
```sh
$ go build
```

## Endpoints:

API endpoint - http://localhost:4000

```sh
GET    /api/user/getAll
POST   /api/user/add
PUT    /api/user/update/{id}
DELETE /api/user/delete/{id}
```

### Get User
This endpoint retrieves a user information.  
Sends a `GET` request to `/api/user/getAll`:
```sh
curl -X GET 'http://localhost:4000/api/user/getAll'
```
Response:
```sh
[
  { 
    "_id":"639e53688390e0e9f0902c26",
    "address":"address2",
    "creation":"19/12/2022",
    "description":"description2",
    "dob":"7/8/2001",
    "name":"Ralph"
   },  
  { 
    "_id":"639e59f7186676441decc92d",
    "address":"address1",
    "creation":"19/12/2022",
    "description":"description1",
    "dob":"7/8/2001",
    "name":"Ralph"
   }, 
  { 
    "_id":"639e5a0c186676441decc92e",
    "address":"address3",
    "creation":"19/12/2022",
    "description":"description3",
    "dob":"8/9/2000",
    "name":"Tom"
   }
]
```
### Create User
This endpoint adds a user to the `users` database.  
Sends a `POST` request to `/api/user/add`:
```sh
curl -X POST 'http://localhost:4000/api/user/add'
```
Response:  
```sh
  { 
    "_id":"000000000000000000000000",
    "name":"Jerry",
    "dob":"19/01/2003",
    "add":"address4",
    "desc":"description4",
    "createdAt":"19/12/2022"
   }

```
### Update User
This endpoint updates the user decription to "Description updated".  
Send a `PUT` request to `/api/user/update/{id}`:
```sh
curl -X PUT 'http://localhost:4000/api/user/update/{id}'
```
Response:
```sh
{
   "id": "<user_id>",
   "name":"Jerry",
   "dob":"19/01/2003",
   "address":"address4",
   "description":"Description updated",             // Description updated
   "creation":"19/12/2022"
}
```

### Delete User
This endpoint deletes the a user from database given the user id.  
Send a `DELETE` request to `/api/user/delete/{id}`:
```sh
curl -X DELETE 'http://localhost:4000/api/user/delete/{id}'
```
Response:
```sh
<user_id>                                           // Returns user id of the deleted record
```
  ### Todos

  - Write unit tests cases
  - User authentication
  - Track login activities
  
### Tech

Open source projects used:

* [Go] - Programing language by Google
* [mux] - Implements a request router and dispatcher in Go
* [MongoDB] - document-based, big community, database
* [ReqBin] - for API testing





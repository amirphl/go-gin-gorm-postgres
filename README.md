# Warm-up exercise with Gin and GORM
 After learning the fundemental of the Golang, I was planning to play with the webframeworks of the Golang. I chose Gin (There are other web frameworks that I haven't compared yet). Unlike Django, Gin doesn't have a built-in ORM, therfore, I selected GORM as the ORM. In this repo, I have implemented a very simple REST service with two entities (models): `User` and `Book`.  The whole of the project is nothing but a bunch of requests which create and modifies the users and the books. If you are interested, you can see my commits from the `Initial commit` then practice yourself.  

## API doc
### Auth
#### register 201
request example (`abc@gmail.com` is a new user):  
```
curl -X POST http://localhost:8080/api/v1/auth/register -d '{"password" : "somepassword1" , "email" : "abc@gmail.com" , "name" : "some-name"}' -H "Content-Type: application/json"
```
response:  
```
{
  "message": "register OK!",
  "errors": null,
  "data": {
    "id": 4,
    "name": "some-name",
    "email": "abc@gmail.com",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJleHAiOjE2NzkwNDY0MzIsImlhdCI6MTY0NzUxMDQzMiwiaXNzIjoiYW1pcnBobC5pciJ9.d18cw4thlYoemwv029wY0cmNCws5YCX1V1_wZzUrfpk"
  }
}
```
#### register 409
request example (`abc@gmail.com` is already registered):  
```
curl -X POST http://localhost:8080/api/v1/auth/register -d '{"password" : "somepassword1" , "email" : "abc@gmail.com" , "name" : "some-name"}' -H "Content-Type: application/json"
```
response:  
```
{
  "message": "Duplicate email",
  "errors": [
    ""
  ],
  "data": null
}
```
#### register 400
request example (`password` is missed):  
```
curl -X POST http://localhost:8080/api/v1/auth/register -d '{"email" : "abc@gmail.com" , "name" : "some-name"}' -H "Content-Type: application/json"
```
response:  
```
{
  "message": "Failed to process register request",
  "errors": [
    "Key: RegisterDTO.Password Error:Field validation for Password failed on the required tag"
  ],
  "data": null
}
```
#### login 200
request example:  
```
curl -X POST http://localhost:8080/api/v1/auth/login -d '{"password" : "somepassword1" , "email" : "abc@gmail.com"}' -H "Content-Type: application/json"
```
response:  
```
{
  "message": "Login OK!",
  "errors": null,
  "data": {
    "id": 4,
    "name": "some-name",
    "email": "abc@gmail.com",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJleHAiOjE2NzkwNDY3OTMsImlhdCI6MTY0NzUxMDc5MywiaXNzIjoiYW1pcnBobC5pciJ9.MdDj3KVYvaeXYKvg3_lKig2eNP_6W07BPKZx_DDa4FQ"
  }
}
```
#### login 401
request example (`password` is incorrect):
```
curl -X POST http://localhost:8080/api/v1/auth/login -d '{"password" : "incorrect-password" , "email" : "abc@gmail.com"}' -H "Content-Type: application/json"
```
response:
```
{
  "message": "Invalid login credential",
  "errors": [
    ""
  ],
  "data": null
}
```
#### login 400
request example (`password` is not provided):
```
curl -X POST http://localhost:8080/api/v1/auth/login -d '{"email" : "abc@gmail.com"}' -H "Content-Type: application/json"
```
response:
```
{
  "message": "Failed to process login request",
  "errors": [
    "Key: LoginDTO.Password Error:Field validation for Password failed on the required tag"
  ],
  "data": null
}
```
### User
Note: the resource `someid` is unused since the real `user-id` is extracted from the JWT sent by the requester in the `Authorization` header.
#### update 204
request example:
```
curl -X PUT http://localhost:8080/api/v1/users/someid -d '{"password" : "somepassword1", "email" : "abc@gmail.com", "name" : "this-is-new-name"}' -H "Content-Type: application/json" -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJleHAiOjE2NzkwNDY3OTMsImlhdCI6MTY0NzUxMDc5MywiaXNzIjoiYW1pcnBobC5pciJ9.MdDj3KVYvaeXYKvg3_lKig2eNP_6W07BPKZx_DDa4FQ'

```
response:
```
```
#### update 401
request example (`JWT` is missed):
```
curl -X PUT http://localhost:8080/api/v1/users/someid -d '{"password" : "somepassword1", "email" : "abc@gmail.com", "name" : "this-is-new-name"}' -H "Content-Type: application/json"
```
response:
```
{
  "message": "No token found",
  "errors": [
    ""
  ],
  "data": null
}
```
#### update 400
request example (`name` is missed):
```
curl -X PUT http://localhost:8080/api/v1/users/someid -d '{"password" : "somepassword1", "email" : "abc@gmail.com"}' -H "Content-Type: application/json" -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJleHAiOjE2NzkwNDY3OTMsImlhdCI6MTY0NzUxMDc5MywiaXNzIjoiYW1pcnBobC5pciJ9.MdDj3KVYvaeXYKvg3_lKig2eNP_6W07BPKZx_DDa4FQ'
```
response:
```
{
  "message": "Failed to process request",
  "errors": [
    "Key: UserUpdateDTO.Name Error:Field validation for Name failed on the required tag"
  ],
  "data": null
}
```
#### retrieve 200
request example:
```
curl -X GET http://localhost:8080/api/v1/users/someid -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJleHAiOjE2NzkwNDY3OTMsImlhdCI6MTY0NzUxMDc5MywiaXNzIjoiYW1pcnBobC5pciJ9.MdDj3KVYvaeXYKvg3_lKig2eNP_6W07BPKZx_DDa4FQ'
```
response:
```
{
  "message": "Get OK!",
  "errors": null,
  "data": {
    "id": 4,
    "name": "this-is-new-name",
    "email": "abc@gmail.com"
  }
}
```
#### retrieve 401
request example (`JWT` is missed):
```
curl -X GET http://localhost:8080/api/v1/users/4
```
response:
```
{
  "message": "No token found",
  "errors": [
    ""
  ],
  "data": null
}
```
### Book
#### create 201
request example:
```
curl -X POST http://localhost:8080/api/v1/books/ -d '{"title" : "Sophie world", "description" : "philosophi"}' -H "Content-Type: application/json" -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJleHAiOjE2NzkwNDY3OTMsImlhdCI6MTY0NzUxMDc5MywiaXNzIjoiYW1pcnBobC5pciJ9.MdDj3KVYvaeXYKvg3_lKig2eNP_6W07BPKZx_DDa4FQ'
```
response:
```
{
  "message": "Book created!",
  "errors": null,
  "data": {
    "id": 4,
    "Title": "Sophie world",
    "Desc": "philosophi",
    "UserID": 4,
    "user": {
      "id": 4,
      "name": "this-is-new-name",
      "email": "abc@gmail.com"
    }
  }
}
```
#### create 401
request example (`JWT` is missed):
```
curl -X POST http://localhost:8080/api/v1/books/ -d '{"title" : "Sophie world", "description" : "philosophi"}' -H "Content-Type: application/json"
```
response:
```
{
  "message": "No token found",
  "errors": [
    ""
  ],
  "data": null
}
```
#### create 400
request example (`description` is missed):
```
curl -X POST http://localhost:8080/api/v1/books/ -d '{"title" : "Sophie world"}' -H "Content-Type: application/json" -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJleHAiOjE2NzkwNDY3OTMsImlhdCI6MTY0NzUxMDc5MywiaXNzIjoiYW1pcnBobC5pciJ9.MdDj3KVYvaeXYKvg3_lKig2eNP_6W07BPKZx_DDa4FQ'
```
response:
```
{
  "message": "Failed to process the request",
  "errors": [
    ""
  ],
  "data": null
}
```
#### update 204
request example:
```
curl -X PUT http://localhost:8080/api/v1/books/4 -d '{"title" : "Sophie world - new title" , "description" : "This is a new desc"}' -H "Content-Type: application/json" -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJleHAiOjE2NzkwNDY3OTMsImlhdCI6MTY0NzUxMDc5MywiaXNzIjoiYW1pcnBobC5pciJ9.MdDj3KVYvaeXYKvg3_lKig2eNP_6W07BPKZx_DDa4FQ'
```
response:
```
```
#### update 401
request example (`JWT` is missed):
```
curl -X PUT http://localhost:8080/api/v1/books/4 -d '{"title" : "Sophie world - new title" , "description" : "This is a new desc"}' -H "Content-Type: application/json"
```
response:
```
{
  "message": "No token found",
  "errors": [
    ""
  ],
  "data": null
}
```
#### update 403
request example (book `10` does not belong to the requester):
```
curl -X PUT http://localhost:8080/api/v1/books/10 -d '{"title" : "Sophie world - new title" , "description" : "This is a new desc"}' -H "Content-Type: application/json" -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJleHAiOjE2NzkwNDY3OTMsImlhdCI6MTY0NzUxMDc5MywiaXNzIjoiYW1pcnBobC5pciJ9.MdDj3KVYvaeXYKvg3_lKig2eNP_6W07BPKZx_DDa4FQ'
```
response:
```
{
  "message": "Action is forbidden",
  "errors": [
    ""
  ],
  "data": null
}
```
#### update 400
request example (`invalid_id` is a `string`, not an `interger`):
```
curl -X PUT http://localhost:8080/api/v1/books/invalid_id -d '{"title" : "Sophie world - new title" , "description" : "This is a new desc"}' -H "Content-Type: application/json" -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJleHAiOjE2NzkwNDY3OTMsImlhdCI6MTY0NzUxMDc5MywiaXNzIjoiYW1pcnBobC5pciJ9.MdDj3KVYvaeXYKvg3_lKig2eNP_6W07BPKZx_DDa4FQ'

```
response:
```
{
  "message": "Failed to process the request",
  "errors": [
    ""
  ],
  "data": null
}
```
#### retrieve 200
request example (everybody can retrieve any book):
```
curl -X GET http://localhost:8080/api/v1/books/1
```
response:
```
{
  "message": "Book found!",
  "errors": null,
  "data": {
    "id": 1,
    "Title": "this is a book title - number 1",
    "Desc": "",
    "UserID": 1,
    "user": {
      "id": 0,
      "name": "",
      "email": ""
    }
  }
}
```
#### retrieve 404
request example (`123` does not exist):
```
curl -X GET http://localhost:8080/api/v1/books/123
```
response:
```
{
  "message": "Book not found",
  "errors": [
    ""
  ],
  "data": null
}
```
####  retrieve 400
request example (`invalid_id` is a `string`, not an `interger`):
```
curl -X GET http://localhost:8080/api/v1/books/invalid_id
```
response:
```
{
  "message": "Failed to process the request",
  "errors": [
    ""
  ],
  "data": null
}
```
#### list 200
request example (everybody can receive the list of all registered books - (no pagination yet)):
```
curl -X GET http://localhost:8080/api/v1/books/
```
response:
```
{
  "message": "Books found!",
  "errors": null,
  "data": [
    {
      "id": 1,
      "Title": "this is a book title - number 1",
      "Desc": "",
      "UserID": 1,
      "user": {
        "id": 1,
        "name": "amirmohammad new haha",
        "email": "amirchanged@gmail.com"
      }
    },
    {
      "id": 3,
      "Title": "this is a book title - number 2 - updateeeeed 2",
      "Desc": "this is the description - number 2",
      "UserID": 1,
      "user": {
        "id": 1,
        "name": "amirmohammad new haha",
        "email": "amirchanged@gmail.com"
      }
    },
    {
      "id": 4,
      "Title": "Sophie world - new title",
      "Desc": "This is a new desc",
      "UserID": 4,
      "user": {
        "id": 4,
        "name": "this-is-new-name",
        "email": "abc@gmail.com"
      }
    }
  ]
}
```
#### delete 200
request example:
```
curl -X DELETE http://localhost:8080/api/v1/books/4 -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJleHAiOjE2NzkwNDY3OTMsImlhdCI6MTY0NzUxMDc5MywiaXNzIjoiYW1pcnBobC5pciJ9.MdDj3KVYvaeXYKvg3_lKig2eNP_6W07BPKZx_DDa4FQ'
```
response:
```
{
  "message": "Book Deleted!",
  "errors": null,
  "data": {
    "id": 4,
    "Title": "Sophie world - new title",
    "Desc": "This is a new desc",
    "UserID": 4,
    "user": {
      "id": 0,
      "name": "",
      "email": ""
    }
  }
}
```
#### delete 401
request example (`JWT` is missed):
```
curl -X DELETE http://localhost:8080/api/v1/books/4
```
response:
```
{
  "message": "No token found",
  "errors": [
    ""
  ],
  "data": null
}
```
#### delete 403
request example (`1` does not belong to the requester):
```
curl -X DELETE http://localhost:8080/api/v1/books/1 -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJleHAiOjE2NzkwNDY3OTMsImlhdCI6MTY0NzUxMDc5MywiaXNzIjoiYW1pcnBobC5pciJ9.MdDj3KVYvaeXYKvg3_lKig2eNP_6W07BPKZx_DDa4FQ'
```
response:
```
{
  "message": "Action is forbidden",
  "errors": [
    ""
  ],
  "data": null
}
```
#### delete 400
request example (`invalid_id` is a `string`, not an `interger`):
```
curl -X DELETE http://localhost:8080/api/v1/books/invalid_id -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJleHAiOjE2NzkwNDY3OTMsImlhdCI6MTY0NzUxMDc5MywiaXNzIjoiYW1pcnBobC5pciJ9.MdDj3KVYvaeXYKvg3_lKig2eNP_6W07BPKZx_DDa4FQ'
```
response:
```
{
  "message": "Failed to process the request",
  "errors": [
    ""
  ],
  "data": null
}
```
## TODO
- token validation 400 to 401
- gitignore
- licence
- gin boilerplate
- clean architecture samples
- github gorm
- poc golang
- binding:"required"
- fix gojwt incompatible
- google: golang crypto
- clean arch of the system
- write tests
- fix api returns password
- sql injection
- who can update a book? write or anyone?
- who can update a user? himself or someone else?
- validation on ID field of BookUpdateDTO and UserUpdateDTO
- partial update of user
- partial update of book
- django vs gin
- django orm vs grom
- github.com/golang-jwt/jwt v3.2.2+incompatible
- test UpdateUser by providing invalid ID in payload
- // TODO inside codes
- link to teutorial
- jwt service: get user refactoring
- remove duplicate codes
- book search: does not exist
- complete README
- test partial updates _test.go

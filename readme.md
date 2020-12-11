RESTFULL AUTH/ACCOUNT MANAGING HTTP-SERVICE
==============================================

**Brief summary:**

+ uses sqlite3 as RDBMS
+ uses gin-gonic anf gorm as a framework and an ORM correspondingly
+ logs are streamed to stdout
+ uses JWT for authentication

**Usage**

+ to run the app, use `make run` command, this will bring the server up on port 8080
+ wil automigrate DB schema at first run. Db file is `sqlite.db` in the project's root

**Request Samples**

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/191d8d14f81921ce7371)

Creating a new account:

````
curl --location --request POST '127.0.0.1:8080/account' \
--header 'Content-Type: application/json' \
--data-raw '{
    "login": "vasya15699",
    "password": "password",
    "firstname": "John",
    "lastname": "Doe",
    "age": 91,
    "email": "johndoe@mail.int"
}'
````

Log into existing account

````
curl --location --request POST '127.0.0.1:8080/account/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "login": "vasya15",
    "password": "password"
}'
````

Log out from your account

````
curl --location --request POST '127.0.0.1:8080/account/logout' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBZ2UiOjIsIkNyZWF0ZWRBdCI6IjIwMjAtMTItMTFUMDM6MTc6MzQuMjY5MTQxNDk5KzAyOjAwIiwiRW1haWwiOiJqb2huZG9lQG1haWwuaW50IiwiRmlyc3ROYW1lIjoiSm9obiIsIkxhc3ROYW1lIjoiRG9lIiwiTG9naW4iOiJ2YXN5YTQiLCJVcGRhdGVkQXQiOiIyMDIwLTEyLTExVDAzOjE3OjM0LjI2OTE0MTQ5OSswMjowMCIsImV4cCI6MTYwNzY1OTQ5MywiaWQiOjIsIm9yaWdfaWF0IjoxNjA3NjU1ODkzfQ.63v0hzdugm4cVJxBNIahEHUsDcat3p65zpvBmU0UjrI'
````

Retrieving a list of existing accounts

````
curl --location --request GET '127.0.0.1:8080/auth/account' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBZ2UiOjIsIkNyZWF0ZWRBdCI6IjIwMjAtMTItMTFUMDM6MTc6MzQuMjY5MTQxNDk5KzAyOjAwIiwiRW1haWwiOiJqb2huZG9lQG1haWwuaW50IiwiRmlyc3ROYW1lIjoiSm9obiIsIkxhc3ROYW1lIjoiRG9lIiwiTG9naW4iOiJ2YXN5YTQiLCJVcGRhdGVkQXQiOiIyMDIwLTEyLTExVDAzOjE3OjM0LjI2OTE0MTQ5OSswMjowMCIsImV4cCI6MTYwNzY1OTQ5MywiaWQiOjIsIm9yaWdfaWF0IjoxNjA3NjU1ODkzfQ.63v0hzdugm4cVJxBNIahEHUsDcat3p65zpvBmU0UjrI'
````

Editing account properties (will work only for your account, will return 401 Unauthorized otherwise)

````
curl --location --request PUT '127.0.0.1:8080/auth/account/3' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBZ2UiOjIsIkNyZWF0ZWRBdCI6IjIwMjAtMTItMTFUMDU6Mjk6MDEuOTUyODg4ODExKzAyOjAwIiwiRW1haWwiOiJqb2huZG9lQG1haWwuaW50IiwiRmlyc3ROYW1lIjoiSm9obiIsIkxhc3ROYW1lIjoiRG9lIiwiTG9naW4iOiJ2YXN5YTE1IiwiVXBkYXRlZEF0IjoiMjAyMC0xMi0xMVQwNToyOTowMS45NTI4ODg4MTErMDI6MDAiLCJleHAiOjE2MDc2NjM3MjEsImlkIjozLCJvcmlnX2lhdCI6MTYwNzY2MDEyMX0.GvhhDGoA3ulHkFEoSBYX-bvI-SiD2fCIy6SdnPOTgs4' \
--header 'Content-Type: application/json' \
--data-raw '{
    "login": "1vasya156",
    "password": "password",
    "firstname": "1John",
    "lastname": "Doe",
    "age": 15,
    "email": "johndoe@mail.int"
}'
````


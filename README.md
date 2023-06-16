<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [TODO-list](#todo-list)
  - [README](#readme)
    - [Instruction for Running the App](#instruction-for-running-the-app)
    - [Instruction for Testing the App](#instruction-for-testing-the-app)
    - [Instruction for Building the App](#instruction-for-building-the-app)
    - [Interface Documentation](#interface-documentation)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# TODO-list

## README

### Instruction for Running the App

Before running the app, please go dir conf and create new file 'app.ini'. 'app.ini' will be the environment file used by this app.

MySQL will be used for database system for this app. Import the given .SQL file to your local database.

Running the app using this command.
```bash
go run .
```

### Instruction for Testing the App

### Instruction for Building the App

```bash
go build main.go

./main
```

### Interface Documentation

REST API: Request and Response

### Login 

#### Request

`POST /login`

    curl -H 'Accept: application/json' -d '{"login_method": "facebook"}' http://localhost:8080/login

Login_method can be "facebook", "gmail" or "github".

#### Response

    {
        "status_code": 200,
        "message": "ok",
        "data": "https://www.facebook.com/v3.2/dialog/oauth?client_id=109661905486499&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fauth%2Ffacebook%2Fcallback&response_type=code&scope=email&state=i44MD32txHqx7JQih5MowtkojnkVVDFAYcc6"
    }

### Add ToDo

#### Request

`POST /api/todo/addTodo`

    curl -H 'Accept: application/json' -H 'Authorization: {token}' -d '{"todo": "todo things"}' http://localhost:8080/api/todo/addTodo


#### Response

    {
        "status_code": 201,
        "message": "added_to-do_item",
        "data": {}
    }

Add ToDo item.

### List ToDo

#### Request

`POST /api/todo/listTodo`

    curl -H 'Accept: application/json' -H 'Authorization: {token}' -d '{"record": 10, "page": 1, "sort": "id", "order": "asc"}' http://localhost:8080/api/todo/listTodo


#### Response

    {
        "status_code": 200,
        "message": "to-do_list",
        "data": {
            "pagination": {
                "record": 10,
                "page": 1,
                "sort": "id",
                "order": "asc",
                "total_record": 2
            },
            "rows": [
                {
                    "id": 1,
                    "item": "to-do things",
                    "marked": 0,
                    "created_by": "jk",
                    "created_at": "2023-06-13 19:42:34"
                }
            ]
        }
    }

List All ToDo item.

### Delete ToDo

#### Request

`POST /api/todo/deleteTodo`

    curl -H 'Accept: application/json' -H 'Authorization: {token}' -d '{"id": 1}' http://localhost:8080/api/todo/deleteTodo


#### Response

    {
        "status_code": 200,
        "message": "deleted_to-do_item",
        "data": {}
    }

Delete an ToDo item.

### Mark ToDo

#### Request

`POST /api/todo/markTodo`

    curl -H 'Accept: application/json' -H 'Authorization: {token}' -d '{"id": 1}' http://localhost:8080/api/todo/markTodo


#### Response

    {
        "status_code": 200,
        "message": "marked_to-do_item",
        "data": {}
    }

Mark or Unmark an ToDo item.

### Logout

#### Request

`POST /logout`

    curl -H 'Accept: application/json' -H 'Authorization: {token}' http://localhost:8080/logout


#### Response

    {
        "status_code": 200,
        "message": "logout_completed",
        "data": {}
    }

Logout user.

# Messaging Service

## Endpoints

### `POST` /login

Request Body

```
{
    "username":"Can",
    "password":"123456"
}

```
### `POST` /register

Request Body

```
{
    "username":"Can",
    "password":"123456"
}

```

### `POST` /send

Request Body

```
{
    "fromUser":"Can",
    "toUser":"Ali",
    "msg":"Hello",
    "date":"26.02.2021",
}

```

### `POST` /block

Request Body

```
{
    "username":"Can",
    "blockedUser":"Ali"
}

```

### `GET` /view

Request Body

```
{
    "username":"Can",
    "password":"123456"
}

```

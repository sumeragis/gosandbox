# ENVIRONMENT

| component      | version |
| -------------- | ------- |
| go             | 1.2     |
| golang-migrate | 4.15    |

# ENDPOINT

/healthz

/user

- GET
- POST
- PATCH
- DELETE

# DEVELOPMENT

Run database

```
$ make up
$ make migrate-up
```

Run application server

```
$ make dev
```

# DELIVERY

Kick GitHub Action "ecr-push"

Update ECS task revision on AWS

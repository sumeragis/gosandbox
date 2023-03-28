# ENVIRONMENT
|component|version|
|---|---|
|go|1.2|
|golang-migrate|4.15|

# ENDPOINT

/healthz

/user 
- GET
- POST

# DEVELOPMENT
Build database
```
$ make up
$ make migrate-up
```
Run application server
```
$ make dev
```

# TODO
- [ ] test
- [ ] lint
- [ ] ci
- [ ] cd

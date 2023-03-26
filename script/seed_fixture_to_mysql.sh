#!/bin/bash

docker-compose exec -T mysql bash -c 'mysql -u docker -pdocker general < /dml/000001_insert_user.sql'
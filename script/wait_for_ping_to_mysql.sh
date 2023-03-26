#!/bin/bash
until docker-compose exec -T mysql bash -c 'mysqladmin ping -h localhost -P 3306 --silent';
do
  echo "Waiting for database connection..."
  sleep 3
done

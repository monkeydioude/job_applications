#!/bin/bash

# Wait for MySQL to be ready
until mysql -u root -e 'SELECT 1'; do
  echo 'waiting for mysql'
  sleep 2
done

# Create a root user that can connect from any host and grant all privileges
# mysql -u root -e "
#   CREATE USER IF NOT EXISTS 'root'@'%' IDENTIFIED BY 'root';
#   GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
#   FLUSH PRIVILEGES;
# "

# Create the initial database
# mysql -u root -e "CREATE DATABASE IF NOT EXISTS core;"

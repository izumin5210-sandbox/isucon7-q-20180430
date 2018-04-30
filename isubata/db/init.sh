#!/bin/bash

DB_DIR=$(cd $(dirname $0) && pwd)
cd $DB_DIR

mysql -uroot -h 127.0.0.1 -e "DROP DATABASE IF EXISTS isubata; CREATE DATABASE isubata;"
mysql -uroot -h 127.0.0.1 isubata < ./isubata.sql
mysql -uroot -h 127.0.0.1 isubata < ./index.sql

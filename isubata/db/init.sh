#!/bin/bash

DB_DIR=$(cd $(dirname $0) && pwd)
cd $DB_DIR

mysql -uroot -h localhost -P 3306 -e "DROP DATABASE IF EXISTS isubata; CREATE DATABASE isubata;"
mysql -uroot -h localhost -P 3306 isubata < ./isubata.sql
mysql -uroot -h localhost -P 3306 isubata < ./index.sql
mysql -uroot -h localhost -P 3306 isubata < ./index2.sql
mysql -uroot -h localhost -P 3306 isubata < ./index3.sql

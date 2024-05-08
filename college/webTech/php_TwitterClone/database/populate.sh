#!/bin/bash

sqlite3 database/gwitter.db < database/schema.sql
sqlite3 database/gwitter.db < database/populate.sql
chmod 777 database/gwitter.db

echo "Success"
# store_oauth-api
oauth api for store in golang

curl --location --request GET 'localhost:8083/oauth/access_token/123' \
--data-raw ''

curl --location --request POST 'localhost:8083/oauth/access_token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "access_token": "abc123",
    "user": 1,
    "client_id": 2,
    "expires": 3
}'

install cassandra. https://cassandra.apache.org/doc/latest/getting_started/index.html 
https://cassandra.apache.org/
https://www.datastax.com/nosql

To run it: cassandra
cqlsh

describe keyspaces;
CREATE KEYSPACE oauth WITH REPLICATION = {'class':'SimpleStrategy', 'replication_factor': 1}
;

USE oauth;
describe tables;

CREATE TABLE access_token(access_token varchar PRIMARY KEY, user_id bigint, client_id bigint, expires bigint);

cqlsh:oauth> select * from access_token where access_token='sdfj';

The primary key is used to distribute the writing along the entire cluster.

https://github.com/gocql/gocql

This project currently use https://github.com/mercadolibre/golang-restclient

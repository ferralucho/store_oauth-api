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

There are different grant types: 
1) client id and secret (client credentials) exchange those for an access token.

{
    "grant_type": "client_credentials",
    "username": "id-123",
    "password": "secret-123",
}

2) grant type password: 

{
    "grant_type": "password",
    "username": "emailaddress@gmail.com",
    "password": "123abc",
}

/users/1?access_token=AbC123

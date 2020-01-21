# oauth2
OAuth2.0 Server

## Run from Docker
* $ docker build -t avosa/oauth2:dev .
* $ docker rm OAUTH2API
* $ docker run -d -p 8086:8086 -v $(pwd)/db/:/db/ --name oauth2 avosa/oauth2:dev
* $ docker logs OAUTH2API
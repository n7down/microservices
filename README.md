# Micro-Services
Exercise - make a couple of services with their accompanying database and have them communicate with each other via sync or async.

## Getting Started
1. Generate the self-signed certificate with `openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout key.pem -out cert.pem` from the root dir

## Todo
- [x] Make a users service
- [ ] Make a products service
- [x] Create a gateway - use JWT for auth
- [ ] Create a basket service and uses both the users and products service
- [ ] Use mosquitto for async calls
- [x] Use grpc for sync calls
- [x] Use UUIDs
- [x] Use IsActive = false when deleting
- [x] Use go modules
- [ ] Add mocked tests
- [ ] Build out activity for a user - can use async calls through mosquitto and sync calls to get activity
- [ ] [swagger](https://github.com/go-swagger/go-swagger)
- [x] Add TLS to grpc
- [x] Set the jwt claims - add user info to the claims - username, firstname, lastname
- [ ] Use [logstash, elastic search, kibana/grafana](https://github.com/deviantony/docker-elk) for api monitoring - send logs and time it take for each request
- [ ] Create permissions for different users - for example setup permissions for a users to only use certain routes
- [ ] Make a set of v2 apis that use [graphql and gin](https://github.com/gin-gonic/gin/issues/1217)
- [ ] Create a stored procedure that check if the user exists and creates it otherwise

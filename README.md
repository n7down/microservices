# Micro-Services
Exercise - make a couple of services with their accompanying database and have them communicate with each other via sync or async.

## Getting Started
1. Generate the self-signed certificate with `openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout key.pem -out cert.pem` from the root dir

## Todo
- [ ] Make a users service
- [ ] Make a products service
- [x] Create a gateway - use JWT for auth
- [ ] Create a basket service and uses both the users and products service
- [ ] Use mosquitto for async calls
- [x] Use grpc for sync calls
- [ ] Use UUIDs
- [ ] Use IsActive = false when deleting
- [x] Use go modules
- [ ] Add mocked tests
- [ ] Figure out how to [build microservices in go](https://www.google.com/search?q=go+microservices+example&oq=go+microservices+example&aqs=chrome.0.69i59j69i64l3j69i60l2.1650j0j7&sourceid=chrome&ie=UTF-8)
- [ ] Build out activity for a user - can use acync calls through mosquitto and sync calls to get activity
- [ ] [swagger](https://github.com/go-swagger/go-swagger)
- [ ] Add authication to grpc
- [ ] Set the jwt claims

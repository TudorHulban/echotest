# echotest

Steps to execute:
1 - Build image:
docker build -t echotest .

2 - Run mongodb:
docker run -d -p 27017:27017 --name mongo mongo

3 - Run image:  
docker run -d -p 1323:1323 echotest

Add a decision using basic auth hard coded credentials:<br/>

curl -X POST http://localhost:1323/api/decisions  -H 'Content-Type: application/json' -d '{"name":"X","amount":100}' -H 'Authorization: Basic am9lOnNlY3JldA=='
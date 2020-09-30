# echotest

To build the application image:
docker build -t echotest .

Execute:  
docker run -d -p 1323:1323 echotest

Add a decision:
curl -X POST http://localhost:1323/api/decisions  -H 'Content-Type: application/json' -d '{"name":"X","amount":100}'
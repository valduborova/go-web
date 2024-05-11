# go-web

Simple web service that can convert a string to base64 and back, and convert unix timestamp to a human readable format.

To start the server, open a terminal (Tools > Terminal) and run:
```
go run cmd/main.go
```
After the server has been started, you can check its functionality by opening a second terminal and running the command:
```
curl localhost:3000/ping
````

You will get pong in response.
# aws-lambda-chat-service
## Building the Functions

There are two functions:
+ chat-message.go
+ chat-connect.go

AWS Lambda expects two runtime version:
GOOS=linux and GOARCH=amd64

So you can build it like this:

`export GOOS=linux GOARCH=amd64 CGO_ENABLED=0 && go build -o ./bin/chat-message/main chat-message.go helpers.go && go build -o ./bin/chat-connect/main chat-connect.go helpers.go`

For this example, the `helpers.go` file is a dependency for both functions.

Zip them and upload to your Lambda Function in AWS

## Using the `client.html` file

Replace the YOUR_WEBSOCKET_URL string with your WebSocket URL.

```javascript
const socket = new WebSocket('YOUR_WEBSOCKET_URL');
```

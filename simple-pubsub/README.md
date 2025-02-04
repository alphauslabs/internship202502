Run server:

```bash
$ go run main.go
```
Replace IP with the VM's IP address.

Subscribe to a topic using Curl:
```bash
$ curl http://{IP}:8080/subscribe?topic=sampletopic
```

Publish a message to a topic using Curl:

```bash
$ curl -X POST http://{IP}:8080/publish \
-H "Content-Type: application/json" \
-d '{"topic":"sampletopic", "message":"Hello, World!"}'
```
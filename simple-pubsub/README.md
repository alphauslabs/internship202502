Run server:

```bash
$ go run main.go
```

Subscribe to a topic using Curl:

```bash
$ curl http://localhost:8080/subscribe?topic=sampletopic
```

Publish a message to a topic using Curl:

```bash
$ curl -X POST http://localhost:8080/publish \
-H "Content-Type: application/json" \
-d '{"topic":"sampletopic", "message":"Hello, World!"}'
```
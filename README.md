# gRPC + OAuth2 Demo (Go)

This is a minimal working example showing how to secure a gRPC service using OAuth2-style bearer tokens.

## 🧩 Structure
```
grpc-oauth/
├── main.go
├── auth.go
└── hello.proto
```

## 🛠️ Setup

1. Install dependencies:
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

2. Generate Go code:
   ```bash
   protoc --go_out=. --go-grpc_out=. hello.proto
   ```

3. Run the server:
   ```bash
   go run main.go server
   ```

4. Run the client (valid token):
   ```bash
   go run main.go client valid-demo-token
   ```

5. Run the client (invalid token):
   ```bash
   go run main.go client bad-token
   ```

✅ Expected output:
```
Response: Hello, OAuth2!
```

❌ Invalid token:
```
rpc error: code = Unauthenticated desc = invalid token
```

---
© 2025 gRPC OAuth Demo

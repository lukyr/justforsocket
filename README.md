# 🔊 Simple WebSocket Broadcaster with Gorilla Mux

This project is a simple Go server that:

- Accepts WebSocket connections using a `streamKey`
- Broadcasts messages via HTTP POST to all connected clients with that `streamKey`

---

## ⚙️ Requirements

- Go 1.18 or higher
- `direnv` for environment variable loading
- `wscat` for testing WebSocket
- `npm` (for installing `wscat`)

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/lukyr/justforsocket.git
cd justforsocket
```

### 2. Install Go dependencies

```bash
go mod download
```

### 3. Make sure direnv is installed and enabled in your shell

```bash
direnv allow
```

### 4. Run the server

```bash
go run main.go
```

### 5. Install wscat

```bash
npm install -g wscat
```

### 6. Test the WebSocket connection

```bash
wscat -c "ws://localhost:9999?streamKey=d4ec6459-ea62cef8-2ca3fbd7-1576860f"

```

### 7. Send a message to the server

```bash
curl -X POST http://localhost:8080/broadcast \
  -H "Content-Type: application/json" \
  -d '{
    "streamKey": "d4ec6459-ea62cef8-2ca3fbd7-1576860f",
    "message": {
      "type": "chat",
      "content": "Hello WebSocket World!"
    }
  }'

```

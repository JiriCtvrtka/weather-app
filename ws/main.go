package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins for simplicity, customize for production
		return true
	},
}

// Connected clients
var clients = make(map[*websocket.Conn]bool)
var clientsMutex = &sync.Mutex{}

// Message history
var messageHistory = []string{}
var historyMutex = &sync.Mutex{}

// Broadcast messages to all clients
func broadcastMessage(messageType int, message []byte) {
	historyMutex.Lock()
	messageHistory = append(messageHistory, string(message))
	historyMutex.Unlock()

	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	for client := range clients {
		err := client.WriteMessage(messageType, message)
		if err != nil {
			fmt.Printf("Error broadcasting to client: %v\n", err)
			client.Close()
			delete(clients, client)
		}
	}
}

// WebSocket handler
func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to upgrade connection: %v\n", err)
		return
	}
	defer conn.Close()

	clientsMutex.Lock()
	clients[conn] = true
	clientsMutex.Unlock()

	// Send message history to the newly connected client
	historyMutex.Lock()
	for _, msg := range messageHistory {
		err = conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			fmt.Printf("Error sending history to client: %v\n", err)
			break
		}
	}
	historyMutex.Unlock()

	fmt.Println("Client connected")

	for {
		// Read message from client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			break
		}

		fmt.Printf("Received: %s\n", message)

		// Broadcast message to all connected clients
		broadcastMessage(messageType, message)
	}

	clientsMutex.Lock()
	delete(clients, conn)
	clientsMutex.Unlock()
}

// Simple HTTP endpoint
func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello, this is a REST API endpoint"}`))
}

// HTML page for WebSocket connection
func htmlHandler(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<head>
	<title>WebSocket Test</title>
	<script>
		let ws;
		function connect() {
			ws = new WebSocket("ws://" + window.location.host + "/ws");

			ws.onmessage = function(event) {
				document.getElementById("messages").innerText += "\n" + event.data;
			};

			ws.onopen = function() {
				console.log("WebSocket connection opened");
			};

			ws.onclose = function() {
				console.log("WebSocket connection closed");
			};
		}

		function sendMessage() {
			const message = document.getElementById("messageInput").value;
			ws.send(message);
		}
	</script>
</head>
<body>
	<h1>WebSocket Test</h1>
	<button onclick="connect()">Connect</button><br><br>
	<input id="messageInput" type="text" placeholder="Enter message">
	<button onclick="sendMessage()">Send</button>
	<pre id="messages"></pre>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("/ws", wsHandler)   // WebSocket endpoint
	http.HandleFunc("/api", apiHandler) // Simple REST API endpoint
	http.HandleFunc("/", htmlHandler)   // HTML page for WebSocket connection

	fmt.Println("Server started on :8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

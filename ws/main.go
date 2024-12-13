package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins for simplicity, customize for production
		return true
	},
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

	fmt.Println("Client connected")

	for {
		// Read message from client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			break
		}

		fmt.Printf("Received: %s\n", message)

		// Echo message back to client
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			fmt.Printf("Error writing message: %v\n", err)
			break
		}
	}
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

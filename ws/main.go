package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"

	"golang.org/x/net/websocket"
)

// Connected clients
var clients = make(map[*websocket.Conn]bool)
var clientsMutex = &sync.Mutex{}

// Message history
var messageHistory = []string{}
var historyMutex = &sync.Mutex{}

// Broadcast messages to all clients
func broadcastMessage(message []byte) {
	historyMutex.Lock()
	messageHistory = append(messageHistory, string(message))
	historyMutex.Unlock()

	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Printf("Error broadcasting to client: %v\n", err)
			client.Close()
			delete(clients, client)
		}
	}
}

// WebSocket handler
func wsHandler(ws *websocket.Conn) {
	// Register client
	clientsMutex.Lock()
	clients[ws] = true
	clientsMutex.Unlock()

	// Send message history to the newly connected client
	historyMutex.Lock()
	for _, msg := range messageHistory {
		err := websocket.Message.Send(ws, msg)
		if err != nil {
			fmt.Printf("Error sending history to client: %v\n", err)
			break
		}
	}
	historyMutex.Unlock()

	fmt.Println("Client connected")

	// Read messages from the WebSocket connection
	for {
		var msg string
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			break
		}

		fmt.Printf("Received: %s\n", msg)

		// Broadcast message to all connected clients
		broadcastMessage([]byte(msg))
	}

	// Deregister client
	clientsMutex.Lock()
	delete(clients, ws)
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
	tmpl, err := template.ParseFiles("html/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error loading template: %v", err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.Handler(wsHandler).ServeHTTP(w, r)
	}) // WebSocket endpoint
	http.HandleFunc("/api", apiHandler) // Simple REST API endpoint
	http.HandleFunc("/", htmlHandler)   // HTML page for WebSocket connection

	fmt.Println("Server started on :8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

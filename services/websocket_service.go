package services

import (
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
)

var upgrader = websocket.Upgrader{
	// Correct field name and initialization
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connections []*websocket.Conn

// WebSocketHandler handles new WebSocket connections and messages
func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the Origin header and print it
	origin := r.Header.Get("Origin")
	fmt.Printf("Origin: %v\n", origin)

	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Connection failed! Error: ", err)
		return
	}

	// Store the connection in the connections slice
	connections = append(connections, conn)

	// Infinite loop to handle WebSocket communication
	for {
		// Read messages from clients
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message: ", err)
			break
		}

		// Broadcast the message to all connected clients
		for _, c := range connections {
			if err := c.WriteMessage(websocket.TextMessage, message); err != nil {
				fmt.Println("Error writing message: ", err)
				break
			}
		}
	}
}

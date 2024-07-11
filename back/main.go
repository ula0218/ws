package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允許所有的請求來源
		return true
	},
}

var clients = make(map[*websocket.Conn]bool) // 連接的客戶端

func main() {
	http.HandleFunc("/chat", handleConnections)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			// 檢查是否是正常的連接關閉引發的錯誤
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("讀取訊息時發生錯誤: %v", err)
			}
			delete(clients, conn)
			break
		}
		msg := string(msgBytes)
		log.Printf("Received: %s", msg)

		// 假設在這裡處理收到的訊息

		// 發送訊息給客戶端
		err = conn.WriteMessage(websocket.TextMessage, []byte("Received: "+msg))
		if err != nil {
			log.Printf("發送訊息時發生錯誤: %v", err)
			delete(clients, conn)
			break
		}
	}
}

// func sendAliveMessages() {
// 	ticker := time.NewTicker(3 * time.Second) // 每隔 3 秒發送一次 alive 訊息
// 	defer ticker.Stop()

// 	for range ticker.C {
// 		for client := range clients {
// 			// 構建 alive 訊息
// 			message := map[string]string{"message": "alive"}
// 			err := client.WriteJSON(message)
// 			if err != nil {
// 				log.Printf("error sending alive message: %v", err)
// 				client.Close()
// 				delete(clients, client)
// 			}
// 		}
// 	}
// }

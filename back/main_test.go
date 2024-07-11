package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
)

func TestWebSocketConnection(t *testing.T) {
	// 啟動測試伺服器
	server := httptest.NewServer(http.HandlerFunc(handleConnections))
	defer server.Close()

	// 提取伺服器的URL並替換為ws
	url := "ws" + server.URL[4:] + "/chat"

	// 建立 WebSocket 連接
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("websocket connection error: %v", err)
	}
	defer ws.Close() // 確保在測試結束時關閉 WebSocket 連接

	// 發送測試訊息
	testMessage := "hello"
	err = ws.WriteMessage(websocket.TextMessage, []byte(testMessage))
	if err != nil {
		ws.Close()
		t.Fatalf("write message error: %v", err)
	}

	// 讀取伺服器回應
	_, message, err := ws.ReadMessage()
	if err != nil {
		ws.Close()
		t.Fatalf("read message error: %v", err)
	}

	if string(message) != "Received: "+testMessage {
		ws.Close()
		t.Errorf("expected message 'Received: %s', got '%s'", testMessage, string(message))
	}
	ws.Close()
}

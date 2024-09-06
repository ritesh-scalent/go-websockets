package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	router := gin.Default()
	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		// scanner := bufio.NewScanner(conn.NetConn())
		for {
			messageType, r, err1 := conn.NextReader()
			if err1 != nil {
				return
			}
			w, _ := conn.NextWriter(messageType)
			io.Copy(w, r)
			w.Close()

		}
		// conn.WriteMessage(websocket.TextMessage, []byte("hello"))
		// time.Sleep(time.Second)
	})
	router.Run(":8080")
}

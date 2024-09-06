package http

import (
	"io"

	"github.com/gin-gonic/gin"
)

func Chat(c *gin.Context) {
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
}

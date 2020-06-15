package api

import (
	"PC/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SocketOrder(c *gin.Context) {

	var (
		wsConn *websocket.Conn
		err    error
		conn   *Connection
		data   []byte
	)

	// 完成ws协议的握手操作
	// Upgrade:websocket
	if wsConn, err = upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}

	if conn, err = InitConnection(wsConn); err != nil {
		goto ERR
	}

	// 启动线程，不断发消息
	go func() {
		var (
			err error
		)
		for {
			if err = conn.WriteMessage([]byte(utils.GetImg() + "LOVE" + utils.GetFullName() + "拼单成功")); err != nil {
				return
			}
			r := rand.Intn(10) + 4
			time.Sleep(time.Duration(r) * time.Second)
		}
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			fmt.Println("哈哈哈ReadMessageSocketOrder")
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			fmt.Println("嘿嘿嘿WriteMessageSocketOrder")
			goto ERR
		}
	}

ERR:
	conn.Close()

}

var (
	wsConn *websocket.Conn
	err    error
	conn   *Connection
	data   []byte
)

func AdminSocketOrder(c *gin.Context) {

	if wsConn, err = upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}

	if conn, err = InitConnection(wsConn); err != nil {
		goto ERR
	}

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}

		fmt.Println(string(data), "---admin--ReadMessage----")
		if err = conn.WriteMessage([]byte("新的data")); err != nil {
			fmt.Println(string(data), "-----admin---WriteMessage-")
			goto ERR
		}
	}

ERR:
	conn.Close()

}

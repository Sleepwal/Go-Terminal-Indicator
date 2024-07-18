package internal

import (
	"github.com/gorilla/websocket"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

// 币安行情信息websocket地址
const (
	wsEndpoint = "wss://fstream.binance.com/stream?streams=btcusdt@depth"
)

func (s *Server) Run() error {
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		return err
	}

	go ReadMessage(conn)

	// fmt.Println(conn)
	select {}
}

package server

import (
	"log"
	"time"

	"github.com/SleepWlaker/GoTerminalIndicator/model"
	"github.com/gorilla/websocket"
	"github.com/nsf/termbox-go"
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
	termbox.Init()

	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		return err
	}

	var (
		ob  = model.NewOrderbook()
		res model.BinanceDepthRespnese
	)

	go func() {
		for {
			if err := conn.ReadJSON(&res); err != nil {
				log.Fatal(err)
			}
			ob.HandleDepthResponse(res.Data)
			// iter := ob.Asks.Iterator(nil, nil)
			// for iter.Next() {
			// 	fmt.Printf("%+v\n", iter.Item())
			// }
			time.Sleep(time.Second * 2)
		}
	}()

	isRunning := true
	go func() {
		time.Sleep(time.Second * 5)
		isRunning = false
	}()

	for isRunning {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		ob.Render(0, 0)
		termbox.Flush()
	}

	return nil
}

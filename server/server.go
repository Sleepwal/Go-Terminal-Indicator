package server

import (
	"log"
	"strconv"

	"github.com/SleepWlaker/GoTerminalIndicator/glabol"
	"github.com/SleepWlaker/GoTerminalIndicator/model"
	"github.com/SleepWlaker/GoTerminalIndicator/term"
	"github.com/gorilla/websocket"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

// 币安行情信息websocket地址
const (
	wsEndpoint     = "wss://fstream.binance.com/stream?streams=btcusdt@depth"
	wsEndpointMark = "wss://fstream.binance.com/stream?streams=btcusdt@markPrice"
)

var (
	WIDTH  = 0
	HEIGHT = 0
)

func (s *Server) Run() error {
	var (
		ob     = model.NewOrderbook()
		result map[string]any
	)

	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		return err
	}

	go func() {
		for {
			if err := conn.ReadJSON(&result); err != nil {
				log.Fatal(err)
			}

			data := result["data"].(map[string]any)
			asks := data["a"].([]any)
			bids := data["b"].([]any)
			ob.HandleDepthResponse(asks, bids)
		}
	}()

	connMark, _, err := websocket.DefaultDialer.Dial(wsEndpointMark, nil)
	if err != nil {
		return err
	}
	go func() {
		for {
			if err := connMark.ReadJSON(&result); err != nil {
				log.Fatal(err)
			}

			glabol.PrevMarkPrice = glabol.CurrMarkPrice
			data := result["data"].(map[string]any)
			priceStr := data["p"].(string)
			glabol.FundingRate = data["r"].(string)
			glabol.CurrMarkPrice, _ = strconv.ParseFloat(priceStr, 64)
		}
	}()

	term.RenderUI(ob)
	return nil
}

package internal

import (
	"log"
	"time"

	"github.com/SleepWlaker/GoTerminalIndicator/internal/model"
	"github.com/gorilla/websocket"
)

func ReadMessage(conn *websocket.Conn) {
	var (
		ob  = model.NewOrderbook()
		res model.BinanceDepthRespnese
	)

	for {
		if err := conn.ReadJSON(&res); err != nil {
			log.Fatal(err)
		}

		// fmt.Println(res)
		ob.HandleDepthResponse(res.Data)
		time.Sleep(time.Second * 2)
	}
}

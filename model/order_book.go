package model

import (
	"fmt"
	"strconv"

	"github.com/SleepWlaker/GoTerminalIndicator/term"
	"github.com/VictorLowther/btree"
	"github.com/nsf/termbox-go"
)

// 订单条目
type OrderbookEntry struct {
	Price  float64 // 价格
	Volume float64 // 数量
}

// 订单
type Orderbook struct {
	Asks *btree.Tree[*OrderbookEntry] // 卖单
	Bids *btree.Tree[*OrderbookEntry] // 买单
}

func byBestBid(a, b *OrderbookEntry) bool {
	return a.Price >= b.Price
}

func byBestAsk(a, b *OrderbookEntry) bool {
	return a.Price < b.Price
}

func NewOrderbook() *Orderbook {
	return &Orderbook{
		Asks: btree.New(byBestAsk),
		Bids: btree.New(byBestBid),
	}
}

func (ob *Orderbook) HandleDepthResponse(result BinanceDepthResult) {
	for _, ask := range result.Asks { // 价格
		price, _ := strconv.ParseFloat(ask[0], 64)
		volume, _ := strconv.ParseFloat(ask[1], 64)
		if volume == 0 {
			continue
		}

		entry := &OrderbookEntry{
			Price:  price,
			Volume: volume,
		}
		ob.Asks.Insert(entry) // 插入
	}

	for _, bid := range result.Bids { // 数量
		price, _ := strconv.ParseFloat(bid[0], 64)
		volume, _ := strconv.ParseFloat(bid[1], 64)
		if volume == 0 {
			continue
		}

		entry := &OrderbookEntry{
			Price:  price,
			Volume: volume,
		}
		ob.Bids.Insert(entry) // 插入
	}
}

func (ob *Orderbook) Render(x, y int) {
	iter := ob.Asks.Iterator(nil, nil)
	i := 0
	for iter.Next() {
		item := iter.Item()
		price := fmt.Sprintf("%.2f", item.Price)
		term.RenderText(x, y+i, price, termbox.ColorGreen)
		i++
	}
}

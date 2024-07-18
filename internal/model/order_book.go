package model

import (
	"fmt"
	"strconv"

	"github.com/VictorLowther/btree"
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
	for _, ask := range result.Asks {
		price, _ := strconv.ParseFloat(ask[0], 64)
		volume, _ := strconv.ParseFloat(ask[1], 64)
		entry := &OrderbookEntry{
			Price:  price,
			Volume: volume,
		}

		fmt.Printf("%+v\n", entry)
	}
}

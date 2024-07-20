package model

import (
	"sort"
	"strconv"
)

// 订单
type Orderbook struct {
	Asks map[float64]float64 // 卖单
	Bids map[float64]float64 // 买单
}

func NewOrderbook() *Orderbook {
	return &Orderbook{
		Asks: make(map[float64]float64),
		Bids: make(map[float64]float64),
	}
}

func (ob *Orderbook) HandleDepthResponse(asks, bids []any) {
	for _, v := range asks { // 价格
		ask := v.([]any)
		price, _ := strconv.ParseFloat(ask[0].(string), 64)
		volume, _ := strconv.ParseFloat(ask[1].(string), 64)
		ob.addAsk(price, volume)
	}

	for _, v := range bids { // 数量
		bid := v.([]any)
		price, _ := strconv.ParseFloat(bid[0].(string), 64)
		volume, _ := strconv.ParseFloat(bid[1].(string), 64)
		ob.addBid(price, volume)
	}
}

func (ob *Orderbook) addAsk(price, volume float64) {
	if volume == 0 {
		delete(ob.Asks, price)
		return
	}
	ob.Asks[price] = volume
}

func (ob *Orderbook) addBid(price, volume float64) {
	if volume == 0 {
		delete(ob.Bids, price)
		return
	}
	ob.Bids[price] = volume
}

func (ob *Orderbook) getAsks() []OrderbookEntry {
	depth := 10
	entries := make(byBestAsk, len(ob.Asks))
	i := 0
	for price, volume := range ob.Asks {
		entries[i] = OrderbookEntry{Price: price, Volume: volume}
		i++
	}

	sort.Sort(entries)
	if len(entries) > depth { // 深度超过depth舍弃
		entries = entries[:depth]
	}
	return entries
}

func (ob *Orderbook) getBids() []OrderbookEntry {
	depth := 10
	entries := make(byBestBid, len(ob.Bids))
	i := 0
	for price, volume := range ob.Bids {
		if volume == 0 {
			continue
		}

		entries[i] = OrderbookEntry{Price: price, Volume: volume}
		i++
	}

	sort.Sort(entries)
	if len(entries) > depth { // 深度超过depth舍弃
		entries = entries[:depth]
	}
	return entries
}

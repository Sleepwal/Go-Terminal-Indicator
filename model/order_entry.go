package model

// 订单条目
type OrderbookEntry struct {
	Price  float64 // 价格
	Volume float64 // 数量
}

type byBestAsk []OrderbookEntry

func (a byBestAsk) Len() int           { return len(a) }
func (a byBestAsk) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byBestAsk) Less(i, j int) bool { return a[i].Price < a[j].Price }

type byBestBid []OrderbookEntry

func (b byBestBid) Len() int           { return len(b) }
func (b byBestBid) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byBestBid) Less(i, j int) bool { return b[i].Price > b[j].Price }

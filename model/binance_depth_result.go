package model

type BinanceDepthResult struct {
	// [price][size(volume)]
	Asks [][]string `json:"a"`
	Bids [][]string `json:"b"`
}

type BinanceDepthRespnese struct {
	Stream string             `json:"stream"`
	Data   BinanceDepthResult `json:"data"`
}

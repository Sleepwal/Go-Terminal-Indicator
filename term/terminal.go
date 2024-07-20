package term

import (
	"fmt"
	"time"

	"github.com/SleepWlaker/GoTerminalIndicator/glabol"
	"github.com/SleepWlaker/GoTerminalIndicator/model"
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func RenderUI(ob *model.Orderbook) {
	termui.Init()

	margin, pheight := 2, 3

	pticker := widgets.NewParagraph()
	pticker.Title = "Binance"
	pticker.Text = "[BTCUSDT](fg:cyan)"
	pticker.SetRect(0, 0, 14, pheight)

	pprice := widgets.NewParagraph()
	pprice.Title = "Market Price"
	ppriceOffset := 14 + 14 + margin + 2
	pprice.SetRect(14+margin, 0, ppriceOffset, pheight)

	pfund := widgets.NewParagraph()
	pfund.Title = "Funding rate"
	pfund.SetRect(ppriceOffset+margin, 0, ppriceOffset+margin+16, pheight)

	tob := widgets.NewTable()
	out := make([][]string, 20)
	for i := 0; i < 20; i++ {
		out[i] = []string{"n/a", "n/a"}
	}
	tob.TextStyle = termui.NewStyle(termui.ColorWhite)
	tob.SetRect(0, pheight+2, 30, 22+pheight+2)
	tob.PaddingBottom = 0
	tob.PaddingTop = 0
	tob.RowSeparator = false
	tob.TextAlignment = termui.AlignCenter

	run := true
	go func() {
		time.Sleep(1 * time.Minute)
		run = false
	}()

	for run {
		if !run {
			return
		}

		var (
			asks = ob.GetAsks()
			bids = ob.GetBids()
		)

		if len(asks) >= 10 {
			for i := 0; i < 10; i++ {
				out[i] = []string{fmt.Sprintf("[%.2f](fg:red)", asks[i].Price), fmt.Sprintf("[%.2f](fg:cyan)", asks[i].Volume)}
			}
		}

		if len(bids) >= 10 {
			for i := 0; i < 10; i++ {
				out[i+10] = []string{fmt.Sprintf("[%.2f](fg:green)", bids[i].Price), fmt.Sprintf("[%.2f](fg:cyan)", bids[i].Volume)}
			}
		}
		tob.Rows = out

		pprice.Text = GetMarketPrice()
		pfund.Text = fmt.Sprintf("[%s](fg:cyan)", glabol.FundingRate)
		termui.Render(pticker, pprice, pfund, tob)
		time.Sleep(time.Millisecond * 20)
	}
}

func GetMarketPrice() string {
	price := fmt.Sprintf("[%s %.2f](fg:green)", glabol.ARROW_UP, glabol.CurrMarkPrice)
	if glabol.PrevMarkPrice > glabol.CurrMarkPrice {
		price = fmt.Sprintf("[%s %.2f](fg:red)", glabol.ARROW_DOWN, glabol.PrevMarkPrice)
	}
	return price
}

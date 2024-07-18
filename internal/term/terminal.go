package term

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func RenderText(x int, y int, msg string, color termbox.Attribute) {
	for _, ch := range msg {
		termbox.SetCell(x, y, ch, color, termbox.ColorDefault)
		w := runewidth.RuneWidth(ch)
		x += w
	}
}

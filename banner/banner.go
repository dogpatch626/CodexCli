package banner

import (
	"github.com/common-nighthawk/go-figure"
)

func Banner() {

	banner := figure.NewColorFigure("Codex", "isometric1", "green", true)
	banner.Print()
}

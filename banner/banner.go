package banner

import (
	"github.com/common-nighthawk/go-figure"
)

func Banner() {

	banner := figure.NewColorFigure("Codex", "banner3-D", "white", true)
	banner.Print()
}

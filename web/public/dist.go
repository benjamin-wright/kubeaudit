package dist

import (
	"embed"
)

//go:embed *
var f embed.FS

func GetWebContent() embed.FS {
	return f
}

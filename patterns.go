package patterns

import "embed"

//go:embed patterns/*.txt
var Patterns embed.FS

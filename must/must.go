package must

import (
	m "github.com/binaryphile/must"
	"time"
)

var (
	TimeParse = m.Must2(time.Parse)
)

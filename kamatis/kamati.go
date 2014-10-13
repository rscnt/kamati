package kamati

import (
	"time"
)

type KamatiProperties struct {
    kamatiesTaskLength  time.Duration
    kamatiesRestLength  time.Duration
    kamatiesBreakLength time.Duration
}

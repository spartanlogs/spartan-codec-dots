package codecs

import (
	"github.com/spartanlogs/spartan/codecs"
	"github.com/spartanlogs/spartan/event"
	"github.com/spartanlogs/spartan/utils"
)

// The DotCodec "dot" converts an event to a single dot.
type DotCodec struct {
	codecs.BaseCodec
}

func init() {
	codecs.Register("dots", newDotCodec)
}

func newDotCodec(options utils.InterfaceMap) (codecs.Codec, error) {
	return &DotCodec{}, nil
}

// Encode event as a dot.
func (c *DotCodec) Encode(e *event.Event) []byte {
	return []byte{'.'}
}

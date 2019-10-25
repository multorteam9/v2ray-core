// +build !confonly

package reverse

import (
	"crypto/rand"
	"io"

	"v2ray.com/core/v4/common/dice"
)

func (c *Control) FillInRandom() {
	randomLength := dice.Roll(64)
	c.Random = make([]byte, randomLength)
	io.ReadFull(rand.Reader, c.Random)
}

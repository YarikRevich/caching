package lru

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"

	. "github.com/franela/goblin"
)

func TestLru(t *testing.T) {
	g := Goblin(t)

	g.Describe("Test lru cache", func() {
		const (
			key      = "test"
			value    = "test"
			capacity = 5
		)

		l := New(capacity)

		g.It("Should return true", func() {

			l.Set(Cell{Key: key, Value: value})

			g.Assert(l.Get(key)).Equal(value)
		})

		g.It(fmt.Sprintf("Should have %d elements", capacity), func() {
			for i := 0; i <= capacity * 2; i++ {
				randInt, err := rand.Int(rand.Reader, big.NewInt(1000))
				if err != nil {
					panic(err)
				}
				l.Set(Cell{Key: randInt.String(), Value: randInt.String()})
			}
			g.Assert(l.QueueLen()).Equal(capacity)
		})

	})
}

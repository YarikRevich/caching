package lru_test

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/YarikRevich/lru/pkg/interfaces"
	"github.com/YarikRevich/lru/pkg/lru"

	. "github.com/franela/goblin"
)

func TestLru(t *testing.T) {
	g := Goblin(t)

	g.Describe("Test lru cache", func() {
		const Key = "test"
		const Value = "test"
		const Capacity = 10

		l := lru.New(Capacity)

		g.It("Should return true", func() {

			l.Set(interfaces.Cell{Key: Key, Value: Value})

			g.Assert(l.Get(Key)).Equal(Value)
		})

		g.It("Should have only 10 elements", func() {
			for i := 0; i <= 20; i++ {
				randInt, err := rand.Int(rand.Reader, big.NewInt(1000))
				if err != nil {
					panic(err)
				}
				l.Set(interfaces.Cell{Key: randInt.String(), Value: randInt.String()})
			}
			g.Assert(l.Len()).Equal(10)
		})
	})
}

package lru_test

import (
	"crypto/rand"
	"fmt"
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
		const Capacity = 1

		l := lru.New(Capacity)

		g.It("Should return true", func() {

			l.Set(interfaces.Cell{Key: Key, Value: Value})

			g.Assert(l.Get(Key)).Equal(Value)
		})

		g.It(fmt.Sprintf("Should have %d elements", Capacity), func() {
			for i := 0; i <= 20; i++ {
				randInt, err := rand.Int(rand.Reader, big.NewInt(1000))
				if err != nil {
					panic(err)
				}
				l.Set(interfaces.Cell{Key: randInt.String(), Value: randInt.String()})
			}
			g.Assert(l.Len()).Equal(Capacity)
		})

		g.It("Should have 0 elements", func() {
			for i := 0; i <= 20; i++ {
				randInt, err := rand.Int(rand.Reader, big.NewInt(1000))
				if err != nil {
					panic(err)
				}
				l.Set(interfaces.Cell{Key: randInt.String(), Value: randInt.String()})
				l.Get(randInt.String())
			}
			g.Assert(l.Len()).Equal(0)
		})


		g.It("Shouldn't remove elements", func() {
			for i := 0; i <= 20; i++ {
				randInt, err := rand.Int(rand.Reader, big.NewInt(1000))
				if err != nil {
					panic(err)
				}
				l.Set(interfaces.Cell{Key: randInt.String(), Value: randInt.String()})
			
			}
			g.Assert(len(l.GetAllWithoutShift())).Equal(Capacity)
			g.Assert(l.Len()).Equal(Capacity)
		})
	})
}

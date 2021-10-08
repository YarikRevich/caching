package zeroshifter

import (
	"crypto/rand"
	"math/big"
	"testing"

	. "github.com/franela/goblin"
)

func TestLru(t *testing.T) {
	g := Goblin(t)

	g.Describe("Test zeroshifter cache", func() {
		const capacity = 10

		z := New(capacity)

		g.It("Checks work of clean method", func(){
			r, err := rand.Int(rand.Reader, big.NewInt(1000))
			g.Assert(err).IsNil()

			z.Add(r)

			g.Assert(len(z.Get())).Eql(1)

			z.Clean()

			g.Assert(len(z.Get())).Eql(0)
		})

		g.AfterEach(func() {
			z.Clean()
		})

		g.It("Should return equal value", func() {
			r, err := rand.Int(rand.Reader, big.NewInt(1000))
			g.Assert(err).IsNil()

			z.Add(r)

			l := z.Get()
			g.Assert(len(l)).Eql(1)
			g.Assert(l).Equal([]interface{}{r})
		})

		g.It("Should filter correctly", func() {
			r, err := rand.Int(rand.Reader, big.NewInt(1000))
			g.Assert(err).IsNil()

			z.Add(r)

			z.Filter(func(i interface{}) bool {
				return false
			})

			l := z.Get()
			g.Assert(len(l)).Eql(0)
		})

		g.It("Battle task", func(){
			for _, v := range []string{"Error", "Warning", "Debug"}{
				z.Add(v)
			}

			g.Assert(len(z.Get())).Eql(3)

			z.Filter(func(i interface{}) bool {
				return i.(string) != "Debug"
			})

			g.Assert(len(z.Get())).Eql(2)

			z.Add("Debug")
			g.Assert(len(z.Get())).Eql(3)
		})
	})

}

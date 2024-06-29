package module2

import (
	"fmt"
	"math/rand"
	"github.com/vivekprm/algo-in-go/module2/sorttest"
	"testing"
)

// To run the benchmark run
func TestBubbleSortInt(t *testing.T) {
	sorttest.TestInt(t, BubbleSortInt)
}

// go_test -bench Bubble -run Bubble
func BenchmarkBubbleSortInt(b *testing.B) {
	 for _, size := range []int{
	 	100, 200, 400, 800, 1600,
	 } {
	 	b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
	 		for n := 0; n < b.N; n++ {
	 			b.StopTimer()
	 			list := rand.Perm(size)
	 			b.StartTimer()
	 			BubbleSortInt(list)
			}
		})
	}
}

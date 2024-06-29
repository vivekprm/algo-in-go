package module2

import (
	"github.com/vivekprm/algo-in-go/module2/sorttest"
	"testing"
)

func TestInsertionSortInt(t *testing.T) {
	sorttest.TestInt(t, InsertionSortInt)
}

func BenchmarkInsertionSortInt(b *testing.B) {
	sorttest.BenchmarkInt(b, InsertionSortInt)
}

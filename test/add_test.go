package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	square := func(n int) int {
		<- time.After(10 * time.Second)
		return n*n
	}
	sum := Add(square(5), 6)
	assert.Equal(t, 31, sum)
}
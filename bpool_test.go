package bpool

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPool(t *testing.T) {
	pool := Pool{}

	a := pool.Get(10)
	require.NotNil(t, a)
	assert.Equal(t, 0, len(a.B))
	assert.Equal(t, 10, cap(a.B))
	a.B = append(a.B, "one"...)
	pool.Put(a)

	b := pool.Get(20)
	require.NotNil(t, b)
	assert.Equal(t, 0, len(b.B))
	assert.Equal(t, 20, cap(b.B))
	b.B = append(b.B, "two"...)
	pool.Put(b)

	c := pool.Get(10)
	require.NotNil(t, c)
	assert.Equal(t, 0, len(c.B))
	assert.True(t, cap(c.B) >= 10)
	assert.True(t, c == a || c == b)

	t.Run("new and capacity=0", func(t *testing.T) {
		pool := Pool{}
		buf := pool.Get(0)
		require.NotNil(t, buf)
		assert.Nil(t, buf.B)
	})
}

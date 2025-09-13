package bufferedqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BufferedQueue(t *testing.T) {
	bq := NewBufferedQueue(5)
	assert.NotNil(t, bq)
	assert.NotNil(t, bq)
	assert.True(t, bq.IsEmpty()) // []
	assert.EqualValues(t, 0, bq.Len())
	ok, err := bq.Push(1) // [1]
	assert.NoError(t, err)
	assert.True(t, ok)
	ok, err = bq.Push(2) // [1,2]
	assert.NoError(t, err)
	assert.True(t, ok)
	ok, err = bq.Push(3) // [1,2,3]
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.False(t, bq.IsEmpty())
	assert.EqualValues(t, 3, bq.Len())
	ok, err = bq.Push(4) // [1,2,3,4]
	assert.NoError(t, err)
	assert.True(t, ok)
	ok, err = bq.Push(5) // [1,2,3,4,5]
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.EqualValues(t, 5, bq.Len())
	assert.False(t, bq.IsEmpty())
	ok, err = bq.Push(6) // [1,2,3,4,5]
	assert.Error(t, err)
	assert.False(t, ok)
	assert.EqualValues(t, 5, bq.Len())
	val, err := bq.Poll() // [2,3,4,5] val=1
	assert.NoError(t, err)
	assert.EqualValues(t, 1, val)
	lend := bq.Len() // [2,3,4,5]
	assert.EqualValues(t, uint(4), lend)
	val, err = bq.Poll() // [3,4,5] val=2
	assert.NoError(t, err)
	assert.EqualValues(t, 2, val)
	assert.EqualValues(t, 3, bq.Len())
	ok, err = bq.Push(6) // [3,4,5,6]
	assert.NoError(t, err)
	assert.True(t, ok)
	assert.EqualValues(t, 4, bq.Len())
	ok, err = bq.Push(7)
	assert.NoError(t, err)
	assert.True(t, ok) // [3,4,5,6,7]
	assert.EqualValues(t, 5, bq.Len())
	ok, err = bq.Push(8)
	assert.Error(t, err)
	assert.False(t, ok) // [3,4,5,6,7]
	val, err = bq.Poll()
	assert.NoError(t, err)
	assert.EqualValues(t, 3, val) // [4,5,6,7]
	assert.EqualValues(t, 4, bq.Len())
	val, err = bq.Poll()
	assert.NoError(t, err)
	assert.EqualValues(t, 4, val) // [5,6,7]
	assert.EqualValues(t, 3, bq.Len())
}

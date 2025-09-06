package bufferedqueue

import (
	"errors"
)

type Queue interface {
	Poll() (uint, error)
	Push(uint) (bool, error)
	IsEmpty() bool
	Len() uint
}

type bufferedQueue struct {
	array []uint
	head  uint
	tail  uint
	size  uint
}

func NewBufferedQueue(n uint) Queue {
	return &bufferedQueue{
		array: make([]uint, n),
		head:  0,
		tail:  0,
		size:  n,
	}
}

func (bq *bufferedQueue) Poll() (uint, error) {
	if bq.IsEmpty() {
		return 0, errors.New("empty queue")
	}

	index := bq.head % bq.size

	value := bq.array[index]
	bq.head = bq.head + 1

	return value, nil
}

func (bq *bufferedQueue) Push(n uint) (bool, error) {
	if bq.Len() == bq.size {
		return false, errors.New("queue full")
	}

	index := bq.tail % bq.size

	bq.array[index] = n

	bq.tail = bq.tail + 1

	return true, nil
}

func (bq *bufferedQueue) IsEmpty() bool {
	return bq.Len() == 0
}

func (bq *bufferedQueue) Len() uint {
	return bq.tail - bq.head
}

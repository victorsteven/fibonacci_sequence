package domain_test

import (
	"fibonacci_sequence/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	fib = domain.NewFibonacciService()
)

func TestService_FibonacciGenerator(t *testing.T) {

	current1 := fib.FibonacciGenerator("current")
	assert.Equal(t, current1, 0)

	next1 := fib.FibonacciGenerator("next")
	assert.Equal(t, next1, 1)

	next2 := fib.FibonacciGenerator("next")
	assert.Equal(t, next2, 1)

	next3 := fib.FibonacciGenerator("next")
	assert.Equal(t, next3, 2)

	previous1 := fib.FibonacciGenerator("previous")
	assert.Equal(t, previous1, 1)

}

func TestPrev(t *testing.T) {

	assert.Equal(t, domain.Prev(5), 3)
	assert.Equal(t, domain.Prev(3), 2)

}

func TestNext(t *testing.T) {

	assert.Equal(t, domain.Next(5), 8)
	assert.Equal(t, domain.Next(8), 13)

}


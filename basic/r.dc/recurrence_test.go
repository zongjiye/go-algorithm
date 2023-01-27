package r_dc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecurrence(t *testing.T) {
	assert.Equal(t, 120, Recurrence(5), "Recurrence error")
	assert.Equal(t, 120, RecurrenceTail(5, 1), "RecurrenceTail error")
	assert.Equal(t, 8, Fib(5), "Fib error")
	assert.Equal(t, 8, FibTail(5, 1, 1), "FibTail error")
}

func BenchmarkRecurrence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Recurrence(20)
	}
}

func BenchmarkRecurrenceTail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecurrenceTail(20, 1)
	}
}

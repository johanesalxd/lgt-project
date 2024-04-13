package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
	t.Run("increment the counter concurrently", func(t *testing.T) {
		target := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(target)

		for i := 0; i < target; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, target)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}

package ge_test

import (
	"github.com/CharLemAznable/ge"
	"sync"
	"sync/atomic"
	"testing"
)

func TestForEach(t *testing.T) {
	wg := &sync.WaitGroup{}
	var sum atomic.Int64

	wg.Add(1)
	actionFn := ge.ActionFunc[int](func(i int) {
		sum.Add(int64(i))
		wg.Done()
	})

	wg.Add(1)
	ch := make(chan int, 1)
	actionCh := ge.ActionChan[int](ch)
	go func() {
		select {
		case i := <-ch:
			sum.Add(int64(i))
			wg.Done()
		}
	}()

	actions := ge.Actions[int]{actionFn, actionCh}
	ge.ForEach([]ge.Action[int]{actions}, 2)
	wg.Wait()
	if sum.Load() != 4 {
		t.Errorf("Expected sum 4, but got %d", sum.Load())
	}
}

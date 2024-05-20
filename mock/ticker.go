package mock

import "time"

type Timer struct {
	ticks int
}

func NewTimer(ticks int) Timer {
	return Timer{
		ticks: ticks,
	}
}

func (t Timer) C() <-chan time.Time {
	c := make(chan time.Time)

	go func() {
		for i := 0; i < t.ticks; i++ {
			c <- time.Now()
		}
		close(c)
	}()

	return c
}

func (t Timer) Stop() {}

package mytimer

import (
	"time"
)

type TimerWrapper struct {
	timer *time.Timer
}

// NewTimerWrapper 创建一个新的TimerWrapper
func NewTimerWrapper(duration time.Duration, callback func()) *TimerWrapper {
	tw := &TimerWrapper{
		timer: time.NewTimer(duration),
	}

	go func() {
		<-tw.timer.C
		callback()
	}()

	return tw
}

// Stop 停止定时器
func (tw *TimerWrapper) Stop() {
	tw.timer.Stop()
}

func (tw *TimerWrapper) Reset(duration time.Duration) {
	tw.timer.Reset(duration)
}

// TickerWrapper 封装了time.Ticker，提供了更友好的接口
type TickerWrapper struct {
	ticker *time.Ticker
}

// NewTickerWrapper 创建一个新的TickerWrapper
func NewTickerWrapper(interval time.Duration, callback func()) *TickerWrapper {
	tw := &TickerWrapper{
		ticker: time.NewTicker(interval),
	}

	go func() {
		for {
			<-tw.ticker.C
			callback()
		}
	}()

	return tw
}

// Stop 停止周期性定时器
func (tw *TickerWrapper) Stop() {
	tw.ticker.Stop()
}

func (tw *TickerWrapper) Reset(duration time.Duration) {
	tw.ticker.Reset(duration)
}
// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package common

import (
	"sync"
	"time"

	"github.com/ying32/govcl/vcl"
)

type Milliseconds int64

type TNotifyEvent func(sender *TTimer)

type TTimer struct {
	interval time.Duration
	ticker   *time.Ticker
	onStop   chan bool
	onTimer  func(sender *TTimer)
	mutex    sync.Mutex
}

func NewTimer() *TTimer {
	return &TTimer{
		interval: time.Duration(1000 * 1e6),
		onStop:   make(chan bool),
	}
}

func (t *TTimer) SetInterval(value Milliseconds) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.interval = time.Duration(value * 1e6)
}

func (t *TTimer) SetEnabled(value bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !value {
		go t.stop()
	} else {
		if t.ticker == nil {
			go t.start()
		}
	}
}

func (t *TTimer) SetOnTimer(fn TNotifyEvent) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.onTimer = fn
}

// start 开始定时器
func (t *TTimer) start() {
	t.ticker = time.NewTicker(t.interval)
	for {
		select {
		case <-t.ticker.C:
			if t.onTimer != nil {
				vcl.ThreadSync(func() {
					t.onTimer(t)
				})
			}
		case <-t.onStop:
			t.ticker.Stop()
			t.ticker = nil
			return
		}
	}
}

// stop 停止定时器
func (t *TTimer) stop() {
	if t.ticker != nil {
		t.onStop <- true
	}
}

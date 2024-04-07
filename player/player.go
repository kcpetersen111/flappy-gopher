package player

import (
	"sync"
	"time"
)

type Player struct {
	Xpos      float64
	Ypos      float64
	Direction bool
	coolDown  bool
	jumpMutex sync.RWMutex
}

func (p *Player) Jump() {
	p.jumpMutex.Lock()
	defer p.jumpMutex.Unlock()
	if p.coolDown {
		return
	}
	p.coolDown = true
	p.Ypos -= 75
	// go off cool down in go routine
	go func() {
		time.Sleep(time.Second)
		p.jumpMutex.Lock()
		defer p.jumpMutex.Unlock()
		p.coolDown = false
	}()
}

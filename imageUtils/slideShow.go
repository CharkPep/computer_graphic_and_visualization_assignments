package imageUtils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"time"
)

type Slideshow struct {
	images       []*canvas.Image
	window       *fyne.Window
	currentIndex int
	delay        time.Duration
	running      bool
	stopChannel  chan struct{}
}

func New(window *fyne.Window, images []*canvas.Image, delay time.Duration) *Slideshow {
	return &Slideshow{
		images:       images,
		window:       window,
		currentIndex: 0,
		delay:        delay,
		running:      false,
		stopChannel:  make(chan struct{}),
	}
}

func modLikePython(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func (s *Slideshow) Start() {
	(*s.window).Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == fyne.KeyLeft {
			(*s).currentIndex = modLikePython((*s).currentIndex-1, len(s.images))
		}
		if k.Name == fyne.KeyRight {
			(*s).currentIndex = ((*s).currentIndex + 1) % len(s.images)
		}
		s.Refresh()
	})

	if !s.running {
		s.running = true
		go func() {
			for {
				select {
				case <-time.After(s.delay):
					(*s).currentIndex = ((*s).currentIndex + 1) % len(s.images)
					s.Refresh()
				case <-s.stopChannel:
					return
				}
			}
		}()
	}
}

func (s *Slideshow) Stop() {
	if s.running {
		s.running = false
		s.stopChannel <- struct{}{}
	}
}

func (s *Slideshow) Next() {
	(*s).currentIndex = ((*s).currentIndex + 1) % len(s.images)
	s.Refresh()
}

func (s *Slideshow) Refresh() {
	(*s.window).SetContent(s.images[s.currentIndex])
}

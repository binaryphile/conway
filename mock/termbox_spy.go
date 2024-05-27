package mock

import (
	iter "github.com/binaryphile/iterator"
	"github.com/binaryphile/slice"
	"github.com/nsf/termbox-go"
	"strings"
	"time"
)

type EventIterator = iter.Iterator[termbox.Event]

type TermboxSpy struct {
	gridSpy   *[][]rune
	nextEvent EventIterator
}

type GridSpyConfig struct {
	Width, Height int
}

type TestEventIteratorConfig struct {
	KeyEventString string
	Step           time.Duration
}

func NewTermboxSpy(gsc GridSpyConfig, kec TestEventIteratorConfig) *TermboxSpy {
	return &TermboxSpy{
		gridSpy:   NewGridSpy(gsc),
		nextEvent: NewTestEventIterator(kec),
	}
}

func (s *TermboxSpy) Clear(_, _ termbox.Attribute) error {
	return nil
}

func (s *TermboxSpy) Close() {}

func (s *TermboxSpy) Flush() error {
	return nil
}

func (s *TermboxSpy) Init() error {
	return nil
}

func (s *TermboxSpy) PollEvent() termbox.Event {
	if s.nextEvent != nil {
		var event termbox.Event
		event, s.nextEvent = s.nextEvent()
		return event
	}

	return termbox.Event{
		Type: termbox.EventKey,
	}
}

func (s *TermboxSpy) String() string {
	if s.gridSpy == nil || len(*s.gridSpy) == 0 {
		return ""
	}

	b := strings.Builder{}
	for j := range (*s.gridSpy)[0] {
		for i := range *s.gridSpy {
			b.WriteRune((*s.gridSpy)[i][j])
		}
		b.WriteRune('\n')
	}

	return b.String()
}

func (s *TermboxSpy) SetCell(x, y int, ch rune, _, _ termbox.Attribute) {
	(*s.gridSpy)[x][y] = ch
}

func (s *TermboxSpy) SetInputMode(mode termbox.InputMode) termbox.InputMode {
	return mode
}

func NewGridSpy(c GridSpyConfig) *[][]rune {
	gridSpy := make([][]rune, c.Width)
	for i := range gridSpy {
		gridSpy[i] = make([]rune, c.Height)
	}

	return &gridSpy
}

func NewTestEventIterator(c TestEventIteratorConfig) EventIterator {
	type RuneSlice = slice.OfTo[rune, termbox.Event]

	events := RuneSlice(c.KeyEventString).Map(EventFromRune)

	i := 0
	var next EventIterator
	next = func() (termbox.Event, EventIterator) {
		time.Sleep(c.Step)

		event := events[i]
		if i == len(events)-1 {
			return event, nil
		}
		i += 1

		return event, next
	}

	return next
}

func EventFromRune(key rune) termbox.Event {
	return termbox.Event{
		Type: termbox.EventKey,
		Ch:   key,
	}
}

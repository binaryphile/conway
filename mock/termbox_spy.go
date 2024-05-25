package mock

import (
	"github.com/binaryphile/conway/ternary"
	iter "github.com/binaryphile/iterator"
	"github.com/nsf/termbox-go"
	"strings"
)

type eventIterator = iter.Iterator[termbox.Event]

type TermboxSpy struct {
	gridSpy   *[][]rune
	nextEvent eventIterator
}

func NewTermboxSpy(events []termbox.Event) *TermboxSpy {
	If := ternary.If[iter.Iterator[termbox.Event]]
	gridSpy := make([][]rune, 0)
	return &TermboxSpy{
		gridSpy:   &gridSpy,
		nextEvent: If(len(events) > 0).Then(iter.FromSlice(events)).Else(nil),
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
	}

	return "\n" + b.String()
}

func (s *TermboxSpy) SetCell(x, y int, ch rune, _, _ termbox.Attribute) {
	(*s.gridSpy)[x][y] = ch
}

func (s *TermboxSpy) SetInputMode(mode termbox.InputMode) termbox.InputMode {
	return mode
}

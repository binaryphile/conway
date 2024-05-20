package mock

import (
	"github.com/binaryphile/conway/iterator"
	"github.com/binaryphile/conway/ternary"
	"github.com/nsf/termbox-go"
	"strings"
)

type eventIterator = iterator.Iter[termbox.Event]

type TermboxSpy struct {
	gridSpy   *[][]rune
	nextEvent eventIterator
}

func NewTermboxSpy(gridSpy *[][]rune, events []termbox.Event) *TermboxSpy {
	If := ternary.If[iterator.Iter[termbox.Event]]
	return &TermboxSpy{
		gridSpy:   gridSpy,
		nextEvent: If(len(events) > 0).Then(iterator.FromSlice(events)).Else(nil),
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

func (s *TermboxSpy) GridString() string {
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

func (s *TermboxSpy) SetCell(_, _ int, _ rune, _, _ termbox.Attribute) {}

func (s *TermboxSpy) SetInputMode(mode termbox.InputMode) termbox.InputMode {
	return mode
}

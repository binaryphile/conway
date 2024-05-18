package mock

import "github.com/nsf/termbox-go"

type TermboxSpy struct {
	methodSpy []string
}

func NewTermboxSpy(methodSpy []string) *TermboxSpy {
	return &TermboxSpy{
		methodSpy: methodSpy,
	}
}

func (s *TermboxSpy) Clear(_, _ termbox.Attribute) error {
	s.methodSpy = append(s.methodSpy, "Clear")
	return nil
}

func (s *TermboxSpy) Close() {
	s.methodSpy = append(s.methodSpy, "Close")
}

func (s *TermboxSpy) Flush() error {
	s.methodSpy = append(s.methodSpy, "Flush")
	return nil
}

func (s *TermboxSpy) Init() error {
	s.methodSpy = append(s.methodSpy, "Init")
	return nil
}

func (s *TermboxSpy) PollEvent() termbox.Event {
	s.methodSpy = append(s.methodSpy, "PollEvent")
	return termbox.Event{
		Type: termbox.EventKey,
	}
}

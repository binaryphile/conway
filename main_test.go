package main

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/binaryphile/conway/mock"
	"github.com/binaryphile/conway/userinterface"
	"github.com/binaryphile/must"
	"github.com/google/go-cmp/cmp"
	"github.com/nsf/termbox-go"
	"testing"
	"time"
)

func Test_run(t *testing.T) {
	spies := make(map[string]*mock.TermboxSpy)

	type fields struct {
		app App
	}

	type args struct {
		initialState string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "basic",
			fields: fields{
				app: NewTestApp(TestAppConfig{
					termbox: NewTestTermboxSpy(spies, "basic"),
					ticks:   5,
				}),
			},
			args: args{
				initialState: heredoc.Doc(`
					#_#
					___
					_#_
				`),
			},
			want: heredoc.Doc(`
				___
				_#_
				___
			`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.app.Run(tt.args.initialState)
			got := spies[tt.name].String()
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("run() mismatch (- expected, + actual):\n%s", diff)
			}
		})
	}
}

func NewTestTermboxSpy(spies map[string]*mock.TermboxSpy, name string) *mock.TermboxSpy {
	spy := mock.NewTermboxSpy(NewKeyEvents("a"))
	spies[name] = spy

	return spy
}

type TestAppConfig struct {
	termbox userinterface.Termbox
	ticks   int
}

func NewTestApp(c TestAppConfig) App {
	return App{
		termbox:       c.termbox,
		newTickerChan: NewTestTickerChanFactory(c.ticks),
	}
}

func NewTestTickerChanFactory(ticks int) TickerChanFactory {
	return func(duration time.Duration) (<-chan time.Time, func()) {
		tickerChan := make(chan time.Time)
		layout := "2006-01-02"
		initialTime := must.Must(time.Parse(layout, "2009-03-01"))

		for i := 0; i < ticks; i++ {
			tickerChan <- initialTime.Add(time.Duration(i) * time.Second)
		}

		return tickerChan, func() {}
	}
}

func NewKeyEvents(s string) []termbox.Event {
	return sliceToEvent(s).Map(NewKeyEvent)
}

func NewKeyEvent(key rune) termbox.Event {
	return termbox.Event{
		Type: termbox.EventKey,
		Ch:   key,
	}
}

type sliceToEvent []rune

func (ts sliceToEvent) Map(fn func(rune) termbox.Event) []termbox.Event {
	results := make([]termbox.Event, len(ts))

	for i, t := range ts {
		results[i] = fn(t)
	}

	return results
}

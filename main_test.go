package main

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/binaryphile/conway/mock"
	"github.com/binaryphile/conway/userinterface"
	"github.com/binaryphile/must"
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

func Test_run(t *testing.T) {
	spies := make(map[string]*mock.TermboxSpy)

	type args struct {
		app          App
		initialState string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				app: NewTestApp(TestAppConfig{
					termbox: NewTestTermboxSpy(spies, "basic"),
					ticks:   5,
				}),
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
			Run(tt.args.app, tt.args.initialState)
			got := spies[tt.name].GridString()
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("run() mismatch (- expected, + actual):\n%s", diff)
			}
		})
	}
}

func NewTestTermboxSpy(spies map[string]*mock.TermboxSpy, name string) *mock.TermboxSpy {
	var gridSpy [][]rune
	spy := mock.NewTermboxSpy(&gridSpy, nil)
	spies[name] = spy

	return spy
}

type TestAppConfig struct {
	termbox userinterface.Termbox
	ticks   int
}

func NewTestApp(c TestAppConfig) App {
	return App{
		termbox:           c.termbox,
		tickerChanFactory: NewTestTickerChanFactory(c.ticks),
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

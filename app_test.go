package main

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/binaryphile/conway/mock"
	"github.com/binaryphile/conway/must"
	ui "github.com/binaryphile/conway/userinterface"
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

func Test_run(t *testing.T) {
	spies := make(map[string]*mock.TermboxSpy)

	type fields struct {
		app App
	}

	type args struct {
		initialState State
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
				app: NewTestApp(
					TermboxSpyConfig{
						spies: spies,
						name:  "basic",
					},
					TestTickerChanFactoryConfig{
						step:  1 * time.Millisecond,
						ticks: 1,
					},
				),
			},
			args: args{
				initialState: StateFromString(heredoc.Doc(`
					#_#
					___
					_#_
				`)),
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

func NewTermboxSpy(c TermboxSpyConfig) *mock.TermboxSpy {
	spy := mock.NewTermboxSpy(
		mock.GridSpyConfig{
			Width:  3,
			Height: 3,
		},
		mock.TestEventIteratorConfig{
			KeyEventString: "a",
			Step:           2 * time.Millisecond,
		},
	)
	c.spies[c.name] = spy

	return spy
}

type TermboxSpyConfig struct {
	name  string
	spies map[string]*mock.TermboxSpy
}

type TestTickerChanFactoryConfig struct {
	step  time.Duration
	ticks int
}

func NewTestApp(tsc TermboxSpyConfig, tcc TestTickerChanFactoryConfig) App {
	return App{
		newTickerChan: NewTestTickerChanFactory(tcc),
		ui:            ui.NewTermboxUI(NewTermboxSpy(tsc)),
	}
}

func NewTestTickerChanFactory(c TestTickerChanFactoryConfig) TickerChanFactory {
	return func(duration time.Duration) (<-chan time.Time, func()) {
		tickerChan := make(chan time.Time, c.ticks)
		layout := "2006-01-02"
		initialTime := must.TimeParse(layout, "2009-03-01")

		go func() {
			for i := 0; i < c.ticks; i++ {
				time.Sleep(c.step)
				tickerChan <- initialTime.Add(time.Duration(i) * time.Second)
			}
		}()

		return tickerChan, func() {
			close(tickerChan)
		}
	}
}

package system_test

import (
	marsrobotcmd "github.com/caevv/marsrobot/cmd/marsrobot"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommand(t *testing.T) {
	var tests = []struct {
		name string
		in   []string
		out  []string
	}{
		{
			"remember scent",
			[]string{
				"1 2",
				"0 0 N",
				"FFF",
				"0 0 N",
				"FFFFFFF",
			},
			[]string{
				"0 3 N LOST",
				"0 2 N",
			},
		},
		{
			"simple RFL",
			[]string{
				"1 1",
				"0 1 E",
				"RLF",
			},
			[]string{
				"1 1 E",
			},
		},
		{
			"outbounds robot",
			[]string{
				"1 1",
				"2 2 E",
				"F",
			},
			[]string{
				"3 2 E LOST",
			},
		},
		{
			"everything",
			[]string{
				"5 3",
				"1 1 E",
				"RFRFRFRF",
				"3 2 N",
				"FRRFLLFFRRFLL",
				"0 3 W",
				"LLFFFLFLFL",
			},
			[]string{
				"1 1 E",
				"3 3 N LOST",
				"2 3 S",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual, err := marsrobotcmd.Execute(tt.in)
			assert.NoError(t, err)
			assert.Equal(t, tt.out, actual)
		})
	}
}

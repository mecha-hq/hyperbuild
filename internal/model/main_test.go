package model_test

import (
	"fmt"
	"testing"
)

func Test_Parsone(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		path    string
		want    []string
		wantErr error
	}{
		{
			desc: "It runs the basic example correctly",
			path: "testdata/basic.yaml",
			want: []string{
				"one runs first, surprise!",
				"two runs after one!",
				"three is broken",
			},
			wantErr: fmt.Errorf("Pipeline failed"),
		},
	}
	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

		})
	}
}

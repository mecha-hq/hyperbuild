package model_test

import (
	"testing"

	"github.com/omissis/hyperbuild/internal/config"
	"github.com/omissis/hyperbuild/internal/model"
)

func Test_Run(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		path    string
		want    []string
		wantErr bool
	}{
		{
			desc: "It runs the basic example correctly",
			path: "testdata/basic.yaml",
			want: []string{
				"one runs first, surprise!\n",
				"two runs after one!\n",
				"three\n",
			},
			wantErr: false,
		},
		{
			desc: "It runs the basic example correctly",
			path: "testdata/error.yaml",
			want: []string{
				"broken!\n",
			},
			wantErr: true,
		},
	}
	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			m, err := config.ParseYAMLFile(tC.path)
			if err != nil {
				t.Fatalf("manifest parsing failed: %v", err)
			}
			got, err := model.Run(m)
			if (err != nil) != tC.wantErr {
				t.Errorf("error mismatched: expected %v, got %v", tC.wantErr, err)
			}
			if len(got) != len(tC.want) {
				t.Errorf("len mismatched: expected %d, got %d", len(tC.want), len(got))
			}

			for i, line := range got {
				if tC.want[i] != line {
					t.Errorf("line mismatched: expected %s, got %s", tC.want[i], line)
				}
			}
		})
	}
}

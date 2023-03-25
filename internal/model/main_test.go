package model_test

import (
	"fmt"
	"testing"

	"github.com/omissis/hyperbuild/internal/config"
	"github.com/omissis/hyperbuild/internal/model"
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
			m, err := config.ParseYAMLFile(tC.path)
			if err != nil {
				t.Fatalf("manifest parsing failed: %v", err)
			}
			got, err := model.Run(m.Steps)
			if err != nil {
				t.Errorf("execution failed: %v", err)
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

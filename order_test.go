package grel_test

import (
	"testing"

	"github.com/andremedeiros/grel"
)

func TestOrder_SQL(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		have grel.Order
		want string
	}{
		"ascending":  {grel.NewOrder(grel.NewColumn("created_at"), grel.Ascending), `"created_at" ASC`},
		"descending": {grel.NewOrder(grel.NewColumn("created_at"), grel.Descending), `"created_at" DESC`},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if have, want := test.have.SQL(), test.want; have != want {
				t.Errorf("have %q, want %q", have, want)
			}
		})
	}
}

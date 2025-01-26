package grel_test

import (
	"testing"

	"github.com/andremedeiros/grel"
)

func TestTable_SQL(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		have grel.Table
		want string
	}{
		"simple": {grel.Table{Name: "users"}, `"users"`},
		"alias":  {grel.Table{Name: "users", Alias: "u"}, `"users" AS "u"`},
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

func BenchmarkTable_SQL(b *testing.B) {
	t := grel.Table{Name: "users"}

	for i := 0; i < b.N; i++ {
		t.SQL()
	}
}

func BenchmarkTable_SQL_WithAlias(b *testing.B) {
	t := grel.Table{Name: "users", Alias: "u"}

	for i := 0; i < b.N; i++ {
		t.SQL()
	}
}

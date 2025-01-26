package grel_test

import (
	"testing"

	"github.com/andremedeiros/grel"
)

func TestPredicate_SQL(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		have grel.Predicate
		want string
	}{
		"eq":  {grel.Eq(grel.NewColumn("id"), grel.NewValue(1)), "\"id\" = $0"},
		"neq": {grel.Neq(grel.NewColumn("id"), grel.NewValue(1)), "\"id\" != $0"},
		"gt":  {grel.Gt(grel.NewColumn("id"), grel.NewValue(1)), "\"id\" > $0"},
		"goe": {grel.Goe(grel.NewColumn("id"), grel.NewValue(1)), "\"id\" >= $0"},
		"lt":  {grel.Lt(grel.NewColumn("id"), grel.NewValue(1)), "\"id\" < $0"},
		"loe": {grel.Loe(grel.NewColumn("id"), grel.NewValue(1)), "\"id\" <= $0"},
		"and": {
			grel.And(
				grel.Eq(grel.NewColumn("id"), grel.NewValue(1)),
				grel.Eq(grel.NewColumn("name"), grel.NewValue("foo")),
			),
			"(\"id\" = $0 AND \"name\" = $0)",
		},
		"or": {
			grel.Or(
				grel.Eq(grel.NewColumn("id"), grel.NewValue(1)),
				grel.Eq(grel.NewColumn("name"), grel.NewValue("foo")),
			),
			"(\"id\" = $0 OR \"name\" = $0)",
		},
		"nested or inside and": {
			grel.And(
				grel.Eq(grel.NewColumn("id"), grel.NewValue(1)),
				grel.Or(
					grel.Eq(grel.NewColumn("name"), grel.NewValue("foo")),
					grel.Eq(grel.NewColumn("name"), grel.NewValue("bar")),
				),
			),
			"(\"id\" = $0 AND (\"name\" = $0 OR \"name\" = $0))",
		},
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

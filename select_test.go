package grel_test

import (
	"testing"

	"github.com/andremedeiros/grel"
)

func TestSelect_SQL(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		have grel.Select
		want string
	}{
		"without fields": {
			grel.NewSelect(grel.Table{Name: "users"}),
			`SELECT * FROM "users"`,
		},

		"order": {
			grel.NewSelect(grel.Table{Name: "users"}).
				OrderBy(grel.NewOrder(grel.NewColumn("created_at"), grel.Descending)),
			`SELECT * FROM "users" ORDER BY "created_at" DESC`,
		},

		"predicate": {
			grel.NewSelect(grel.Table{Name: "users"}).
				Where(grel.Eq(grel.NewColumn("users.id"), grel.NewValue(1))),
			`SELECT * FROM "users" WHERE "users"."id" = $1`,
		},

		"join": {
			grel.NewSelect(grel.Table{Name: "users"}).
				Join(grel.Table{Name: "posts"}, grel.InnerJoin, grel.Eq(grel.NewColumn("user_id"), grel.NewColumn("id"))),
			`SELECT * FROM "users" INNER JOIN "posts" ON "user_id" = "id"`,
		},

		"join with predicate": {
			grel.NewSelect(grel.Table{Name: "users"}).
				Join(grel.Table{Name: "posts"}, grel.InnerJoin, grel.Eq(grel.NewColumn("user_id"), grel.NewValue(1))),
			`SELECT * FROM "users" INNER JOIN "posts" ON "user_id" = $1`,
		},

		"predicate and join with predicate": {
			grel.NewSelect(grel.Table{Name: "users"}).
				Join(grel.Table{Name: "posts"}, grel.InnerJoin, grel.Eq(grel.NewColumn("user_id"), grel.NewValue(1))).
				Where(grel.Eq(grel.NewColumn("id"), grel.NewValue(1))),
			`SELECT * FROM "users" INNER JOIN "posts" ON "user_id" = $1 WHERE "id" = $2`,
		},

		"duplicate columns": {
			grel.NewSelect(grel.Table{Name: "users"}, grel.NewColumn("id")).
				Column(grel.NewColumn("users.id")),
			`SELECT "users"."id" FROM "users"`,
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

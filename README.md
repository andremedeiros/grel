# Golang Relational Expression Language

## Warnings

* This is a blatant rip-off of the [AREL](https://github.com/rails/rails/blob/main/activerecord/lib/arel.rb) library from Rails.
* This is a work in progress and is not ready for production use.
* This is a pet project and I'm not sure how far I'll take it.
* The API is not stable and **will** change.

## Example

```go
package main

import (
	"fmt"

	"github.com/andremedeiros/grel"
)

func main() {

	sql := grel.NewSelect(grel.Table{Name: "users"}).
		Join(grel.Table{Name: "posts"}, grel.InnerJoin, grel.Eq(grel.NewColumn("posts.user_id"), grel.NewColumn("users.id"))).
		Where(grel.Eq(grel.NewColumn("users.name"), grel.NewValue("andremedeiros")))

	fmt.Println(query.SQL())
}
```

```
$ go run main.go
SELECT * FROM "users" INNER JOIN "posts" ON "posts"."user_id" = "users"."id" WHERE "users"."name" = $1
```

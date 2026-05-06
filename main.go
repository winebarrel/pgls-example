// Package main is a small playground for exercising pgls inside .go
// files. Place the cursor inside any of the backtick-quoted SQL strings
// below and try completion (Ctrl-Space), hover, and the diagnostics
// panel.
//
// pgls treats a string literal as SQL when either:
//   - a "// language=sql" / "// language=postgresql" marker comment
//     sits on the line directly above, or
//   - the literal is passed to a recognised SQL method. The default
//     set covers database/sql (Query, QueryRow, Exec, Prepare, and
//     their *Context variants); extend it via sqlFunctions in
//     .pgls.json. Plain SQL-looking strings outside either path are
//     intentionally ignored.
package main

import (
	"database/sql"
	"fmt"
)

func main() {
	var db *sql.DB

	// Marker form. Hover over `email` to see "users.email varchar".
	// Type "u." after `FROM users u` to see only that table's columns.
	// After "WHERE" with both users and orders in scope, duplicate
	// column names ("id") appear as qualified entries ("u.id" / "o.id").
	// language=sql
	q1 := `
		SELECT u.id, u.email, o.total
		FROM users u
		JOIN orders o ON u.id = o.user_id
		WHERE u.email IS NOT NULL
		ORDER BY o.placed_at DESC
		LIMIT 50
	`

	// CTEs are recognised as virtual tables — `active_users` doesn't
	// false-flag as an unknown table. Column-level validation inside
	// the CTE body is not done in v1.
	// language=sql
	q2 := `
		WITH active_users AS (
			SELECT id, email FROM users WHERE created_at > now() - interval '7 days'
		)
		SELECT * FROM active_users
	`

	// Function-call form — no marker comment needed because db.Query
	// is in the default sqlFunctions set. Try Cmd+click on "users" or
	// "u.email" to jump into schema/users.sql.
	rows, _ := db.Query(`
		SELECT u.id, u.email
		FROM users u
		WHERE u.id = $1
	`, 1)
	_ = rows

	// Without a marker AND outside a recognised function call, pgls
	// ignores the string completely. Add "// language=sql" above, or
	// pass it to db.Query/db.Exec/etc., to opt in.
	notSQL := `SELECT * FROM users`

	// Uncomment to see diagnostics fire (each becomes SQL via the
	// db.Query call, no marker required):
	//
	//   _, _ = db.Query(` + "`SELECT * FROM producsts`" + `)             // unknown table "producsts"
	//   _, _ = db.Query(` + "`SELECT u.bogus FROM users u`" + `)         // column "bogus" not in table "users"
	//   _, _ = db.Query(` + "`SELECT bad_alias.id FROM users u`" + `)    // unknown table or alias "bad_alias"

	fmt.Println(q1, q2, notSQL)
}

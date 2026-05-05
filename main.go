// Package main is a small playground for exercising pgls inside .go
// files. Place the cursor inside any of the backtick-quoted SQL strings
// below and try completion (Ctrl-Space), hover, and the diagnostics
// panel.
//
// pgls only fires on string literals preceded by a JetBrains-style
// marker comment such as "// language=sql" or "// language=postgresql"
// on the line directly above. Plain SQL-looking strings without a
// marker are intentionally ignored to avoid false positives.
package main

import "fmt"

func main() {
	// Hover over `email` to see "users.email varchar". Type "u." after
	// FROM users u to see only that table's columns. After "WHERE" with
	// both users and orders in scope, duplicate column names ("id")
	// appear as qualified entries ("u.id" / "o.id").
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

	// Without a marker comment pgls ignores the string completely. Add
	// "// language=sql" on the line above to opt in.
	notSQL := `SELECT * FROM users`

	// Uncomment and add a marker comment above to see diagnostics fire:
	//
	//   // language=sql
	//   q3 := ` + "`SELECT * FROM producsts`" + `             // unknown table "producsts"
	//   q4 := ` + "`SELECT u.bogus FROM users u`" + `         // column "bogus" not in table "users"
	//   q5 := ` + "`SELECT bad_alias.id FROM users u`" + `    // unknown table or alias "bad_alias"

	fmt.Println(q1, q2, notSQL)
}

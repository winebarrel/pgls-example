# pgls-example

A small reproducible playground for [pgls](https://github.com/winebarrel/pgls)
— exercises completion, hover, and diagnostics in both `.sql` files and
SQL string literals inside `.go` files.

## What's here

```
.pgls.json         -- Tells pgls to load DDL from ./schema (auto-discovered)
schema/
  users.sql        -- CREATE TABLE definitions
  orders.sql
example.sql        -- Standalone SQL file to play with
main.go            -- Backtick SQL strings inside Go source
```

## Setup

1. Install the pgls server binary:

   ```sh
   go install github.com/winebarrel/pgls@latest
   ```

2. Install an editor client:

   - VSCode: <https://github.com/winebarrel/pgls-vscode>
   - Neovim / Vim / Helix: see the [pgls README](https://github.com/winebarrel/pgls#editor-setup)

3. Open this folder in your editor — pgls picks up `.pgls.json`
   automatically, no per-editor `schemaDir` configuration needed.

## Try it

Open `example.sql` or `main.go` and:

- Place the cursor after `FROM ` → completion offers `users`, `orders`.
- After `SELECT ` (with `FROM users u JOIN orders o`) → completion offers
  columns scoped to those two tables, and the duplicate `id` shows as
  `u.id` / `o.id`.
- After `u.` → only the columns of `users`.
- Hover over `email` → `users.email varchar`.
- Hover over `users` → a markdown table of the columns.
- Cmd+click on a table or `u.email` → opens the corresponding row in
  `schema/users.sql`.
- Uncomment the diagnostic-fire lines at the bottom of either file to
  see `pgls` flag the typos.

`main.go` shows two ways pgls picks up SQL strings:

- **Marker form** (`q1`, `q2`): a `// language=sql` comment on the
  line directly above flags the string as SQL.
- **Function-call form** (`db.Query(...)`): no comment needed —
  pgls's default `sqlFunctions` set covers `database/sql`'s
  `Query` / `Exec` / `Prepare` and their `*Context` variants.

Strings outside both paths (`notSQL` in the example) are
intentionally ignored.

## License

MIT

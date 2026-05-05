# pgls-example

A small reproducible playground for [pgls](https://github.com/winebarrel/pgls)
— exercises completion, hover, and diagnostics in both `.sql` files and
SQL string literals inside `.go` files.

## What's here

```
schema/
  users.sql        -- DDL fed to pgls via initializationOptions.schemaDir
  orders.sql
example.sql        -- Standalone SQL file to play with
main.go            -- Backtick SQL strings inside Go source
```

## Setup

1. Install the pgls server binary:

   ```sh
   go install github.com/winebarrel/pgls@latest
   ```

2. Install a pgls editor client. Examples:

   - VSCode: <https://github.com/winebarrel/pgls-vscode>
   - Neovim built-in LSP: see the pgls README

3. Tell the editor where the schema lives. For VSCode add a workspace
   setting at `.vscode/settings.json` (kept out of git intentionally):

   ```json
   {
     "pgls.schemaDir": "schema"
   }
   ```

   For Neovim point `init_options.schemaDir = "schema"` at the same path.

## Try it

Open `example.sql` or `main.go` and:

- Place the cursor after `FROM ` → completion offers `users`, `orders`.
- After `SELECT ` (with `FROM users u JOIN orders o`) → completion offers
  columns scoped to those two tables, and the duplicate `id` shows as
  `u.id` / `o.id`.
- After `u.` → only the columns of `users`.
- Hover over `email` → `users.email varchar`.
- Hover over `users` → a markdown table of the columns.
- Uncomment the diagnostic-fire lines at the bottom of either file to
  see `pgls` flag the typos.

## License

MIT

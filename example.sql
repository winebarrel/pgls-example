-- Open this file with pgls running and try completion / hover.
-- All the same features that work inside .go string literals also
-- work in a plain .sql file.

-- After "FROM " expect tables only; after "SELECT " or "WHERE "
-- expect columns from the FROM-list; after "u." expect users columns.
SELECT u.id, u.email, o.total
FROM users u
JOIN orders o ON u.id = o.user_id
WHERE u.email IS NOT NULL
ORDER BY o.placed_at DESC;

-- CTE: `active_users` resolves as a virtual table.
WITH active_users AS (
    SELECT id, email FROM users WHERE created_at > now() - interval '7 days'
)
SELECT * FROM active_users;

-- Uncomment to see diagnostics:
-- SELECT * FROM producsts;                  -- unknown table
-- SELECT u.bogus FROM users u;              -- column not in table
-- SELECT bad_alias.id FROM users u;         -- unknown qualifier

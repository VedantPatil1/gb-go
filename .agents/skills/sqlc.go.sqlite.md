# SQLc for Go + SQLite Code Generation

## Installation
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.31.1
```

## Project Structure (sqlc.yaml at root)
```
.
├── db/
│   ├── schema.sql          # Schema definition or migrations
│   └── queries.sql         # SQL query definitions
├── models/                 # Generated Go code
└── sqlc.yaml              # Configuration file
```

## Basic Configuration

**sqlc.yaml:**
```yaml
version: "2"
sql:
  - engine: "sqlite"
    schema: "./db/schema.sql"
    queries: "./db/queries.sql"
    gen:
      go:
        package: dbmodels
        out: ./models
```

## Query Annotations

1. **:one** - Returns single struct (SELECT * LIMIT 1)
```sql
-- name: GetUserById :one
SELECT id, name FROM users WHERE id = ? LIMIT 1;
```

2. **:many** - Returns []struct
```sql
-- name: FindBooksByTags :many
SELECT id, title, views FROM books WHERE tags IN (?) ORDER BY title;
```

3. **INSERT + RETURNING (:one)** - Create and get result
```sql
-- name: CreateUser :exec
INSERT INTO users (username, email) VALUES (?, ?);
```

4. **:noresult** - Execute statement, no return value
```sql
-- name: DeleteUser :noresult
DELETE FROM users WHERE id = ?;
```

5. **UPDATE + RETURNING (:one)** - Recommended for updates
```sql
-- name: UpdateBook :exec
UPDATE books SET title = ?, views = views + 1
WHERE id = ? RETURNING *;
```

## Generated Output: models/
```
models/
├── db.go          # Database interfaces (TxQuerier, Querier)
├── models.go      # Auto-generated Go structs from schema
└── queries.sql.go # Type-safe query functions matching your SQL
```

## Generated Code Example

**From INSERT + RETURNING:**
```go
// queries.sql.go - auto-generated
type User struct {
    Id        int64    // Database field
    Username  string
    Email     sql.NullString
    CreatedAt time.Time
}

func CreateUser(ctx context.Context, db DBTX, arg *CreateUserParams) (CreateUserResult, error) {
    return quoter.CreateUser(ctx, arg)
}
```

**Usage in your application:**
```go
result, err := db.CreateUser(context.Background(), &dbmodels.CreateUserParams{
    Username: "john_doe",
    Email:    sql.NullString{String: "john@example.com", Valid: true},
})
id := result.Id           // Generated PK available!
```

## Running Code Generation

```bash
sqlc generate

# Custom config location (yaml must be in current dir):
sqlc generate --file=./queries.sql
```

## Key Benefits

| Manual SQL | SQLc Generated |
|------------|----------------|
| Type mismatches possible | ✅ Compile-time verified |
| Manual struct updates needed | ✅ Auto-syncs with schema |
| No autocomplete for params | ✅ Full IDE support |
| Prone to SQL injection risks | ✅ Parameterized by default |

## Environment Configuration

```yaml
gen: go:
  package: dbmodels
  out: ./models
  # Output pointer types for NULL columns
  emit_pointers_for_null_types: true

  # Add JSON tags
  emit_json_tags: true

  # Rename camelCase fields to PascalCase
  rename:
    user_name: UserName

  # Use prepared statements
  emit_prepared_queries: true

  # Build directive (optional)
  build_tags: ["sqlc-generated"]
```

## Connection Setup (your internal/db/layer)

```go
// Manual connection + use generated interface
import "myproject/models"

func GetDatabase() models.Querier {
    // Using modernc.org/sqlite driver (no CGO, fast)
    db, err := sql.Open("sqlite3", "./data.db")
    if err != nil { /* handle */ }

    // Apply pragmas for performance
    db.Exec(`PRAGMA journal_mode=WAL`)
    db.Exec(`PRAGMA busy_timeout=5000`)

    return models.NewQuerier(db)  // Type-safe interface!
}
```

## Testing Generated Code

```go
func TestCreateUser(t *testing.T) {
    result := NewQuerier(testDB).InsertUser(
        context.Background(),
        UserParams{
            Username: "test",
            Email:    "test@example.com",
        }
    )

    // Verify created record
    var user models.User
    if err := db.SelectOne(&user, `SELECT * FROM users WHERE id = ?`, result.Id); err != nil {
        t.Fatal(err)
    }
}
```

## Migration Workflow

```bash
# Step 1: Update schema in db/schema.sql or create new migration
cat > db/migrations/002_create_posts.sql <<SQL
CREATE TABLE posts (id INTEGER, title TEXT);
SQL

# Step 2: Add queries to db/queries.sql
# Append your new function definitions

# Step 3: Regenerate
sqlc generate

# Commit everything together
git add models/ && git commit -m "regenerated Go code from schema"
```

## SQL Injection Safety

sqlc parameterizes ALL queries automatically. You write:
```sql
-- SQL injection safe
SELECT * FROM users WHERE email = ?;
```
Go:
```go
queter.GetUserByEmail(ctx, user@example.com)  // No escaping needed!
```

## Version Info

- **Current**: v1.31.1
- **Go Support**: ✅ Stable
- **SQLite Support**: ⚠️ Beta (fully functional)

## References

- Docs: https://docs.sqlc.dev/en/stable/
- GitHub: https://github.com/sqlc-dev/sqlc
- Go docs: https://pkg.go.dev/github.com/sqlc-dev/sqlc/gen-go/internal

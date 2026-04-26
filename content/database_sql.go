package content

func init() {
	Register(&Package{
		Name:       "database/sql",
		ImportPath: "database/sql",
		Category:   "Database",
		Summary:    "Generic SQL interface. You supply a driver via blank import (e.g. _ \"github.com/lib/pq\"). Manages a connection pool for you.",
		Sections: []Section{
			{
				Title: "Open and ping",
				Examples: []Example{
					{Title: "Open (does not connect)", Code: `import _ "github.com/mattn/go-sqlite3"

db, err := sql.Open("sqlite3", "app.db")
if err != nil { log.Fatal(err) }
defer db.Close()

if err := db.PingContext(ctx); err != nil { log.Fatal(err) }`},
					{Title: "Pool tuning", Code: `db.SetMaxOpenConns(25)
db.SetMaxIdleConns(25)
db.SetConnMaxIdleTime(5 * time.Minute)
db.SetConnMaxLifetime(30 * time.Minute)`},
				},
			},
			{
				Title: "Query",
				Examples: []Example{
					{Title: "QueryRowContext — single row", Code: `var name string
var age int
err := db.QueryRowContext(ctx, "SELECT name, age FROM users WHERE id = ?", id).
    Scan(&name, &age)
if err == sql.ErrNoRows {
    // not found
}`},
					{Title: "QueryContext — many rows", Code: `rows, err := db.QueryContext(ctx, "SELECT id, name FROM users")
if err != nil { return err }
defer rows.Close()
for rows.Next() {
    var id int
    var name string
    if err := rows.Scan(&id, &name); err != nil { return err }
}
return rows.Err()`},
				},
			},
			{
				Title: "Exec",
				Examples: []Example{
					{Title: "Insert / Update / Delete", Code: `res, err := db.ExecContext(ctx, "INSERT INTO users(name) VALUES(?)", name)
if err != nil { return err }
id, _ := res.LastInsertId()
n, _ := res.RowsAffected()`},
				},
			},
			{
				Title: "Transactions",
				Examples: []Example{
					{Title: "BeginTx + rollback on error", Code: `tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
if err != nil { return err }
defer tx.Rollback() // no-op if Commit succeeds

if _, err := tx.ExecContext(ctx, "UPDATE ..."); err != nil { return err }
if _, err := tx.ExecContext(ctx, "INSERT ..."); err != nil { return err }
return tx.Commit()`},
				},
			},
			{
				Title: "Prepared statements",
				Examples: []Example{
					{Title: "Prepare once, execute many", Code: `stmt, err := db.PrepareContext(ctx, "SELECT name FROM users WHERE id = ?")
if err != nil { return err }
defer stmt.Close()
for _, id := range ids {
    var name string
    stmt.QueryRowContext(ctx, id).Scan(&name)
}`},
				},
			},
			{
				Title: "NULL handling",
				Examples: []Example{
					{Title: "sql.Null* types", Code: `var nm sql.NullString
db.QueryRowContext(ctx, "SELECT middle_name FROM users WHERE id=?", id).Scan(&nm)
if nm.Valid { fmt.Println(nm.String) }`},
				},
			},
		},
	})
}

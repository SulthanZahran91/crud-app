package repository

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoRepository struct {
	db *sql.DB
}

func New(dbPath string) (*TodoRepository, error) {
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return nil, fmt.Errorf("create db dir: %w", err)
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	if err := migrate(db); err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}
	return &TodoRepository{db: db}, nil
}

func migrate(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id        INTEGER PRIMARY KEY AUTOINCREMENT,
		title     TEXT    NOT NULL,
		completed INTEGER NOT NULL DEFAULT 0
	)`)
	return err
}

func (r *TodoRepository) GetAll() ([]Todo, error) {
	rows, err := r.db.Query(`SELECT id, title, completed FROM todos ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		var completed int
		if err := rows.Scan(&t.ID, &t.Title, &completed); err != nil {
			return nil, err
		}
		t.Completed = completed == 1
		todos = append(todos, t)
	}
	if todos == nil {
		todos = []Todo{}
	}
	return todos, rows.Err()
}

func (r *TodoRepository) Create(title string) (Todo, error) {
	res, err := r.db.Exec(`INSERT INTO todos (title, completed) VALUES (?, 0)`, title)
	if err != nil {
		return Todo{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return Todo{}, err
	}
	return Todo{ID: int(id), Title: title, Completed: false}, nil
}

func (r *TodoRepository) Update(id int, title string, completed bool) (Todo, error) {
	comp := 0
	if completed {
		comp = 1
	}
	res, err := r.db.Exec(`UPDATE todos SET title = ?, completed = ? WHERE id = ?`, title, comp, id)
	if err != nil {
		return Todo{}, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return Todo{}, err
	}
	if n == 0 {
		return Todo{}, fmt.Errorf("not found")
	}
	return Todo{ID: id, Title: title, Completed: completed}, nil
}

func (r *TodoRepository) Delete(id int) error {
	res, err := r.db.Exec(`DELETE FROM todos WHERE id = ?`, id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("not found")
	}
	return nil
}

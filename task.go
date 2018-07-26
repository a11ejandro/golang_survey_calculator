package main

import (
  "fmt"
  "database/sql"
)

type Task struct {
  page int
  perPage int
}

func (task *Task)Find(id sql.NullInt64) (*Task) {
  row := db.QueryRow("SELECT page, per_page FROM tasks WHERE id = $1", id)

  t := new(Task)
  err := row.Scan(&t.page, &t.perPage)
    if err == sql.ErrNoRows {
        fmt.Println("task not found")
    } else if err != nil {
        fmt.Println("Task scan error:", err)
    }

  if err != nil {
    fmt.Println(err)
  }

  return t
}

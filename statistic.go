package main

import (
  "fmt"
  "database/sql"
)

type Statistic struct {
  taskId sql.NullInt64
  handlerType string
  collectionSize int
  duration float64
}

func (statistic *Statistic) Save() {
  _, err := db.Exec("INSERT INTO statistics (task_id, handler_type, collection_size, duration) VALUES ($1, $2, $3, $4)",
    statistic.taskId, statistic.handlerType, statistic.collectionSize, statistic.duration)

  if err != nil {
    fmt.Println("Statistic create error", err)
  }
}

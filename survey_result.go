package main

import (
  "fmt"
  "database/sql"
)

type SurveyResult struct {
  value float64
  taskId sql.NullInt64
}

func (surveyResult *SurveyResult) Page(page, perPage int) []*SurveyResult {
  offset := 0 
  if (page > 0) {
    offset = (page - 1) * perPage
  }

  rows, err := db.Query("SELECT survey_results.value, survey_results.task_id FROM survey_results LIMIT $1 OFFSET $2", perPage, offset)

  if err != nil {
    fmt.Println("DB error: ", err)
  }
  defer rows.Close()

  surveyResults := make([]*SurveyResult, 0)
  for rows.Next() {
      sr := new(SurveyResult)
      err := rows.Scan(&sr.value, &sr.taskId)
      if err != nil {
          fmt.Println("Survey Result Scan Error ", err)
      }
      surveyResults = append(surveyResults, sr)
  }

  return surveyResults
}

func (surveyResult *SurveyResult) Save() {
  _, err := db.Exec("INSERT INTO survey_results (value, task_id) VALUES ($1, $2)", surveyResult.value, surveyResult.taskId)

  if err != nil {
    fmt.Println("Result creation error ", err)
  }
}

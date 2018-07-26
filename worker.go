package main

import (
  "os"
  "fmt"
  "time"
  "database/sql"
  "github.com/benmanns/goworker"
)

func registerWorker() {
  redisUrl := os.Getenv("REDIS_URL")

  workerSettings := goworker.WorkerSettings{
    URI:            redisUrl,
  	Connections:    100,
  	Queues:         []string{"go_survey_calculator"},
  	Concurrency:    2,
  	Namespace:      "",
    Interval:       5.0,
  }
  goworker.SetSettings(workerSettings)
  goworker.Register("GoSurveyCalculator", surveyCalculatorWorker)
}

func surveyCalculatorWorker(queue string, args ...interface{}) error {
  startTime := time.Now()

  taskId := sql.NullInt64{Int64: int64(args[0].(float64)), Valid: true}
  task := new(Task).Find(taskId)

  surveyResults := new(SurveyResult).Page(task.page, task.perPage)
  avg := avgForSurvey(surveyResults)

  // create new survey result which stores the result value of current survey
  newResult := new(SurveyResult)
  newResult.value = avg
  newResult.taskId = taskId
  newResult.Save()
  endTime := time.Now()
  duration := endTime.Sub(startTime)

  // write statistic to db
  statistic := new(Statistic)
  statistic.duration = duration.Seconds()
  statistic.taskId = taskId
  statistic.handlerType = "go_worker"
  statistic.collectionSize = task.perPage

  statistic.Save()
  fmt.Println("Completed task", taskId.Int64)
  return nil
}

func avgForSurvey(surveyResults []*SurveyResult) (float64) {
  min, max, sum := 0.0, 0.0, 0.0

  for _, element := range surveyResults {
      if element.value < min {
        min = element.value
      }

      if element.value > max {
        max = element.value
      }

      sum = sum + element.value
  }

  avg := sum / float64(len(surveyResults))

  return (min + max + avg) / 3
}

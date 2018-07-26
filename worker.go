package main

import (
  "fmt"
  "os"
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
  taskId := int64(args[0].(float64))
  fmt.Println("Recieved task with id ", taskId)

  return nil
}

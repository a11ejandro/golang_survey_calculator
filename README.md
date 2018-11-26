# Golang survey worker.

This is a source code of Go worker, described in article [Comparing the performance of RoR and Go applications in the background processing - Part 2](https://medium.com/@alexander.potrakhov/comparing-the-performance-of-ror-and-go-workers-in-the-background-processing-6b16f6f1bdf6).

The application needs [API](https://github.com/a11ejandro/rails_survey_calculator) for interaction with tasks and statistics, set it up first if you haven't.

## Steps to run
1) Get [Go](https://golang.org/doc/install#install)
2) Clone this repo to $GOPATH:/src/
3) Create and fill in the .env in the root of project. The required keys are POSTGRES_USERNAME, POSTGRES_PASSWORD, POSTGRES_HOST, POSTGRES_DATABASE, POSTGRES_PORT, REDIS_URL. Note, they should point the DB and Redis used in [API](https://github.com/a11ejandro/rails_survey_calculator)
4) Run *go build -o golang_survey_calculator.bin*
5) Run *./golang_survey_calculator.bin*

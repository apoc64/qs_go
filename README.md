# qs_go

This app provides a Go based backend for a calorie tracking app. The frontend is deployed at: http://efficient-secretary.surge.sh/. This backend is deployed at: https://steve-qs-go.herokuapp.com/. The frontend repo is: https://github.com/apoc64/qs_frontend. Please note, that the frontend will require changes to the baseURL variable in order to point to a different backend location.

## Go Version

* Go 1.10.3

## Setup

This app uses Gorilla/mux and Postgres using lib/pq. If you need them, run:

``` $ go get github.com/gorilla/mux github.com/lib/pq ```

This app uses dep for package management. If you need it, run:

```
$ brew install dep
$ brew upgrade dep
```

Then, from inside the root folder for this project, run:

``` $ dep init ```

For a blog post related to deploying this app, visit https://medium.com/@sschwedt/deploying-a-go-app-to-heroku-78284e601232

### Database Setup

To run this app locally you must create a Postgres database entitled qs_go. The app should connect to it and run migrations automatically the first time the app, or test suite is run. Local development and testing share a database, and its data is erased at the start of each test.

When running in production, the app will connect to a database from a DATABASE_URL environment variable. This is currently set up for Heroku.

## Running the app locally

To run the server locally without the tests, run:

``` $ go run $(ls -1 *.go | grep -v _test.go) ```

This line is also included in comments at the bottom of server.go

## Running the tests

To run the test suite, run:

``` $ go test ```

## Endpoints

This app includes the following endpoints:

* GET /api/v1/foods/ - Gets all foods

* POST /api/v1/foods/ - Adds a food to the database, with the body of the request in the format: {"food":{"name":"pizza","calories":"400"}}

* /api/v1/foods/:id/ - Gets the food with the specified id

* /api/v1/meals - Gets all meals with their foods

* /api/v1/meals/:id/foods - Gets all foods for the meal with the specified id

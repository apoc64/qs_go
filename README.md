# qs_go

This app provides a Go based backend for a calorie tracking app

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

To run this app locally you must create a Postgres database entitled qs_go

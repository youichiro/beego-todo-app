# Beego Todo App
Todo app Using Beego (Golang), Vue.js, DataBase (MySQL)


## Install

In order to install Beego, execute:

```
go get -u github.com/beego/bee
go get -u github.com/astaxie/beego
```


## Setup

Place this repository in `$GOPATH/src` and set up your database in `main.go`.
In order to create database, execute:

```
mysql -u root -p -e 'CREATE DATABASE beego_todo'
```

## Start app

In order to run server, just execute:

```
bee run
```

Will start our Beego server on port 8080.
Go to localhost:8080

![screenshot](https://github.com/youichiro/beego-todo-app/blob/master/static/img/screenshot.png)

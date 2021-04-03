
## Go-Migraty 
~~migrate, mighty gotta it?~~
 Go-Migraty is a simple way to implement migrations to help you on your development.
Currently working with mysql.

---

* [Install](#install)
* [Features](#features)
* [Examples](#examples)

---
## Install 
```sh
go get -u github.com/boladissimo/go-migraty
```
## Features

 - Runs migrations ~~wow~~
 - Checks if the table already exists before running its script
 

## Examples
The migration files have to be a sql `create table script` , the file name has to be in the format `tableName.sql` e.g. `person.sql`.

to run the migrations, you only need to instantiate a `migraty.Runner` giving `*sql.DB` and the path to the migrations scripts, then run the runner  function `Migrate`
e.g.
```go
migraty := migraty.New(GetDB(), "/scripts/migrations")
migraty.Migrate()
```
    

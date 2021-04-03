package migraty

import (
	"database/sql"
	"fmt"
	"io/ioutil"
)

type Runner interface {
	//Migrate runs the migrations configured in the runner
	Migrate()
}

type runner struct {
	db             *sql.DB
	migrationsPath string
}

//New return an migraty runner instance
func New(db *sql.DB, migrationsPath string) Runner {
	return runner{db: db, migrationsPath: migrationsPath}
}

func (r runner) Migrate() {
	tables := getTableNames(r.migrationsPath)
	dbName := r.getDBName()

	logInfo(tables, dbName)

	for _, table := range tables {
		if !r.tableExists(dbName, table) {
			logInfo(fmt.Sprintf("running %s migration", table))
			r.db.Exec(getMigrationScript(table, r.migrationsPath))
		}
	}
}

//getTableNames return all tables names given path
func getTableNames(migrationsPath string) (tables []string) {
	files, err := ioutil.ReadDir(migrationsPath)
	if err != nil {
		logError(err)
	}

	for _, file := range files {
		fileFullName := file.Name()
		tableName := fileFullName[:len(fileFullName)-4]
		tables = append(tables, tableName)
		logInfo(tableName)
	}

	return
}

//tableExists check if the given table is present on the database
func (r runner) tableExists(dbname, table string) bool {
	result, err := r.db.Query("SELECT * FROM information_schema.tables WHERE table_schema = ? AND table_name = ? LIMIT 1", dbname, table)
	if err != nil {
		logError(err)
	}
	return result.Next()
}

//getMigrationScript returns the migrations script given its table name and path
func getMigrationScript(table, migrationsPath string) string {
	script, err := ioutil.ReadFile(fmt.Sprintf("%s%s.sql", migrationsPath, table))
	if err != nil {
		logError(err)
	}
	return string(script)
}

//getDBName returns the database name from the runners sql.DB
func (r runner) getDBName() (dbName string) {
	stmt := r.db.QueryRow("SELECT DATABASE()")
	err := stmt.Scan(&dbName)

	if err != nil {
		logError(err)
	}
	return
}

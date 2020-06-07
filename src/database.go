package main

import (
	"database/sql"
	"fmt"

	//"fmt"
	"log"

	_ "github.com/lib/pq"
)

type databaseConnection struct{
	conn *sql.DB
}

const (
	host     = "localhost"
	port     = "54320"
	user     = "dbuser"
	password = "secret"
	dbname   = "he"
)

func NewDatabaseConnection() databaseConnection {
	//connStr := "host=localhost port=54320 user=dbuser dbname=he sslmode=disable"

	connStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping() // check connection
	if err != nil {
		panic(err)
	}

	log.Println("connected to DB!")
	return databaseConnection{db}
}

func (db *databaseConnection) Close() {
	err := db.conn.Close()

	if err != nil {
		log.Fatal(err)
	}
}

func (db *databaseConnection) AddHe(he HE) {
	sqlStatement := fmt.Sprintf(
		"insert into HE (HELinkUuid, HeUuid, fname, lname, file, status)" +
			" values ('%s','%s','%s','%s','%s','%s')", he.HELinkUuid, he.HeUuid, he.Student.Firstname, he.Student.Lastname,
			he.File.Text, he.Status)

	_, err := db.conn.Exec(sqlStatement)

	if err != nil {
		panic(err)
	}
}

func (db *databaseConnection) ExistsHe(heUuid string) bool {
	sqlStatement := fmt.Sprintf("select HELinkUuid from HELink where HELinkUuid = '%s'", heUuid)

	row := db.conn.QueryRow(sqlStatement)

	switch err := row.Scan(&heUuid); err {
	case sql.ErrNoRows:
		return false
	case nil:
		return true
	default:
		panic(err)
	}
}

func (db *databaseConnection) AddHelink(hl HELink) {
	sqlStatement := fmt.Sprintf("insert into HELink (HELinkUuid) values ('%s')", hl.HELinkUuid)

	_, err := db.conn.Exec(sqlStatement)

	if err != nil {
		panic(err)
	}
}

func (db *databaseConnection) ExistsHelink(hl HELink) bool {
	sqlStatement := fmt.Sprintf("select HELinkUuid from HELink where HELinkUuid = '%s'", hl.HELinkUuid)

	row := db.conn.QueryRow(sqlStatement)

	switch err := row.Scan(&hl.HELinkUuid); err {
	case sql.ErrNoRows:
		return false
	case nil:
		return true
	default:
		panic(err)
	}
}

func (db *databaseConnection) GetFile(heUuid string) string {
	var file string
	sqlStatement := fmt.Sprintf("select file from HE where HeUuid = '%s'", heUuid)

	row := db.conn.QueryRow(sqlStatement)

	switch err := row.Scan(&file); err {
	//case sql.ErrNoRows:
	//	return false
	case nil:
		return file
	default:
		panic(err)
	}
}

func (db *databaseConnection) SetFile(heUuid string, file File) {
	sqlStatement := "UPDATE HE SET file = $2 WHERE heuuid = $1;"

	_, err := db.conn.Exec(sqlStatement, heUuid, file.Text)

	if err != nil {
		panic(err)
	}
}



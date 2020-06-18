package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type databaseConnection struct{
	conn *sql.DB
}

func NewDatabaseConnection(host string, port int, user string, password string, dbname string) (databaseConnection, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return databaseConnection{}, err
	}

	err = db.Ping() // check connection
	if err != nil {
		return databaseConnection{}, err
	}

	log.Printf("Successfully connected to database: %s:%d | db: %s\n | user: %s", host, port, dbname, user)

	return databaseConnection{db}, nil
}

func (db *databaseConnection) Close() error { // func not necessary
	return db.conn.Close()
}

func (db *databaseConnection) AddHe(he Submission) error {
	sqlStatement := fmt.Sprintf(
		"insert into HE (HELinkUuid, HeUuid, fname, lname, file, status)" +
			" values ('%s','%s','%s','%s','%s','%s')", he.Link.Uuid, he.Uuid, he.Student.Firstname, he.Student.Lastname,
			he.File.Text, he.File.Status)

	_, err := db.conn.Exec(sqlStatement)

	return err
}

func (db *databaseConnection) ExistsHe(heUuid string) (bool, error) {
	sqlStatement := fmt.Sprintf("select heuuid from HE where heuuid = '%s'", heUuid)

	row := db.conn.QueryRow(sqlStatement)

	switch err := row.Scan(&heUuid); err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}

func (db *databaseConnection) AddHelink(helink string) error {
	sqlStatement := fmt.Sprintf("insert into HELink (HELinkUuid) values ('%s')", helink)

	_, err := db.conn.Exec(sqlStatement)

	return err
}

func (db *databaseConnection) ExistsHelink(helink string) (bool, error) {
	sqlStatement := fmt.Sprintf("select HELinkUuid from HELink where HELinkUuid = '%s'", helink)

	row := db.conn.QueryRow(sqlStatement)

	switch err := row.Scan(&helink); err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, nil
	}
}

func (db *databaseConnection) GetFile(heUuid string) (string, error) {
	var file string
	sqlStatement := fmt.Sprintf("select file from HE where HeUuid = '%s'", heUuid)

	row := db.conn.QueryRow(sqlStatement)

	switch err := row.Scan(&file); err {
	case nil:
		return file, nil
	default:
		return "", err
	}
}

func (db *databaseConnection) SetFile(heUuid string, text string) error {
	sqlStatement := "UPDATE HE SET file = $2 WHERE heuuid = $1;"

	_, err := db.conn.Exec(sqlStatement, heUuid, text)

	return err
}



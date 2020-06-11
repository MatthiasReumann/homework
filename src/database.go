package main

import (
	"database/sql"
	"fmt"

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

	return databaseConnection{db}, nil
}

func (db *databaseConnection) Close() error { // func not necessary
	return db.conn.Close()
}

func (db *databaseConnection) AddSubmission(he Submission) error {
	sqlStatement := fmt.Sprintf(
		"insert into Submission (LinkUuid, SubmissionUuid, fname, lname)" +
			" values ('%s','%s','%s','%s')", he.Link.Uuid, he.Uuid, he.Student.Firstname, he.Student.Lastname)
	_, err := db.conn.Exec(sqlStatement)

	if err != nil {
		return err
	}

	sqlStatement = "insert into file (SubmissionUuid, Text, status) values ($1, $2, $3)"
	_, err = db.conn.Exec(sqlStatement, he.Uuid, he.File.Text, he.File.Status)

	return err
}

func (db *databaseConnection) ExistsSubmission(Uuid string) (bool, error) {
	sqlStatement := fmt.Sprintf("select SubmissionUuid from Submission where SubmissionUuid = '%s'", Uuid)

	row := db.conn.QueryRow(sqlStatement)

	switch err := row.Scan(&Uuid); err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}

func (db *databaseConnection) AddLink(linkUuid string) error {
	sqlStatement := fmt.Sprintf("insert into Link (HELinkUuid) values ('%s')", linkUuid)

	_, err := db.conn.Exec(sqlStatement)

	return err
}

func (db *databaseConnection) ExistsLink(linkUuid string) (bool, error) {
	sqlStatement := fmt.Sprintf("select HELinkUuid from Link where HELinkUuid = '%s'", linkUuid)

	row := db.conn.QueryRow(sqlStatement)

	switch err := row.Scan(&linkUuid); err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, nil
	}
}

func (db *databaseConnection) GetFile(submissionUuid string) (File, error) {
	var file File
	sqlStatement := fmt.Sprintf("select Text, status from File where SubmissionUuid = '%s'", submissionUuid)

	row := db.conn.QueryRow(sqlStatement)

	switch err := row.Scan(&file.Text, &file.Status); err {
	case nil:
		return file, nil
	default:
		return File{}, err
	}
}

func (db *databaseConnection) SetFile(submissionUuid string, file File) error {
	sqlStatement := "UPDATE File SET Text = $1, status = $2 WHERE SubmissionUuid = $3;"

	_, err := db.conn.Exec(sqlStatement, file.Text, file.Status, submissionUuid)

	return err
}



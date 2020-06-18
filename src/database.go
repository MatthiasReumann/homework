package main

import (
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"log"
	_ "github.com/lib/pq"
)

type databaseConnection struct {
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

func (db *databaseConnection) Close() error {
	return db.conn.Close()
}

func (db *databaseConnection) GetSubmissions(linkUuid string) ([]string, error) {
	var subs = make([]string,0)
	sqlStatement, _, _ := goqu.From("submission").Select("submissionuuid").Where(
		goqu.C("linkuuid").Eq(linkUuid),
	).ToSQL()

	rows, err := db.conn.Query(sqlStatement)

	if err != nil {
		return []string{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var sub string
		err = rows.Scan(&sub)
		if err != nil {
			return []string{}, err
		}
		subs = append(subs, sub)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		return []string{}, err
	}

	return subs, nil
}

func (db *databaseConnection) AddSubmission(he Submission) error {
	ds := goqu.Insert("submission").Rows(
		goqu.Record{"linkuuid": he.Link.Uuid, "submissionuuid": he.Uuid,
			"fname": he.Student.Firstname, "lname": he.Student.Lastname},
	)

	sqlStatement, _, _ := ds.ToSQL()
	_, err := db.conn.Exec(sqlStatement)

	if err != nil {
		return err
	}

	ds = goqu.Insert("file").Rows(
		goqu.Record{"submissionuuid": he.Uuid, "text": he.File.Text,
			"status": he.File.Status},
	)

	sqlStatement, _, _ = ds.ToSQL()
	_, err = db.conn.Exec(sqlStatement)

	return err
}

func (db *databaseConnection) ExistsSubmission(Uuid string) (bool, error) {
	sqlStatement, _, _ := goqu.From("submission").Select("submissionuuid").Where(
		goqu.C("submissionuuid").Eq(Uuid),
	).ToSQL()

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

func (db *databaseConnection) AddLink(linkUuid string, task []byte) error {
	// goqu can't insert binary data

	sqlStatement := "INSERT INTO \"link\" (\"helinkuuid\", \"task\") VALUES ($1, $2)"

	_, err := db.conn.Exec(sqlStatement, linkUuid, task) // no sql injection possible

	return err
}

func (db *databaseConnection) ExistsLink(linkUuid string) (bool, error) {
	sqlStatement, _, _ := goqu.From("link").Select("helinkuuid").Where(
		goqu.C("helinkuuid").Eq(linkUuid),
	).ToSQL()

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
	sqlStatement, _, _ := goqu.From("file").Select("text", "status").Where(
		goqu.C("submissionuuid").Eq(submissionUuid),
	).ToSQL()

	row := db.conn.QueryRow(sqlStatement)

	switch err := row.Scan(&file.Text, &file.Status); err {
	case nil:
		return file, nil
	default:
		return File{}, err
	}
}

func (db *databaseConnection) SetFile(submissionUuid string, file File) error {
	sqlStatement, _, _ := goqu.From("file").Where(goqu.C("submissionuuid").Eq(submissionUuid)).Update().Set(
		goqu.Record{"text": file.Text, "status": file.Status},
	).ToSQL()

	_, err := db.conn.Exec(sqlStatement)

	return err
}

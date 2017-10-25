package mysql_status

import (
	"database/sql"
	"io"
	"log"

	"github.com/adzeitor/status"
	_ "github.com/go-sql-driver/mysql"
)

type CheckResult struct {
	Name     string
	Response string
	OK       bool
	Error    string
}

type Checker struct {
	url   string
	Name  string
	Query string
	DB    *sql.DB
}

func New(name string, url string, query string) *Checker {
	db, err := sql.Open("mysql", url)
	// FIXME: return err
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(1)
	return &Checker{
		url:   url,
		Name:  name,
		DB:    db,
		Query: query,
	}
}

func (c *Checker) Check() (ok bool, res []status.CheckResult, err error) {
	ok = true

	err = c.DB.Ping()
	if err != nil {
		ok = false
		res = append(res, CheckResult{Name: c.Name, Error: err.Error()})
		return
	}

	var rows *sql.Rows
	rows, err = c.DB.Query(c.Query)
	if err != nil {
		ok = false
		res = append(res, CheckResult{Name: c.Name, Error: err.Error()})
		return
	}

	defer rows.Close()
	for rows.Next() {
		if rows.Err() != nil {
			err = rows.Err()
			ok = false
			res = append(res, CheckResult{Name: c.Name, Error: err.Error()})
			return
		}
	}
	return
}

func (result CheckResult) Render(w io.Writer) error {

	return defaultTemplate.Execute(w, result)
}

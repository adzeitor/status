package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/adzeitor/status"
	"github.com/adzeitor/status/plugins/http_status"
	"github.com/adzeitor/status/plugins/mysql_status"
)

type customStatus struct{}
type customResult struct {
	OK bool
}

func (s customStatus) Check() (bool, []status.CheckResult, error) {
	res := customResult{OK: rand.Intn(1000) > 300}
	return false, []status.CheckResult{res}, nil
}

func (result customResult) Render(w io.Writer) error {
	if result.OK {
		w.Write([]byte(`<li class="list-group-item">OK</li>`))
	} else {
		w.Write([]byte(`<li class="list-group-item list-group-item-danger">FAIL</li>`))
	}
	return nil
}

func main() {
	statusPage := status.New()
	// http
	statusPage.AddWith(http_status.New([]string{
		"https://httpbin.org/status/418",
		"https://httpbin.org/status/200"}),
		status.Config{Interval: 60 * time.Second})

	// mysql
	statusPage.AddWith(mysql_status.New("database.example.com",
		os.Getenv("DATABASE_URL"),
		"SELECT * FROM table LIMIT 1"),
		status.Config{Interval: time.Second})

	statusPage.Add(customStatus{})
	statusPage.Run()

	log.Println("served http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", statusPage))
}

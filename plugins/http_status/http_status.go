package http_status

import (
	"io"
	"net/http"

	"github.com/adzeitor/status"
)

type HTTPCheckResult struct {
	URL   string
	Code  int
	Error string
	Body  string
	OK    bool
}

type HTTPChecker struct {
	urls []string
}

func New(urls []string) *HTTPChecker {
	return &HTTPChecker{
		urls: urls,
	}
}

func (c *HTTPChecker) Check() (ok bool, res []status.CheckResult, err error) {
	var resp *http.Response

	ok = true

	for _, url := range c.urls {
		resp, err = http.Head(url)
		if err != nil {
			res = append(res, HTTPCheckResult{
				URL:   url,
				Error: err.Error(),
				OK:    false,
			})
			ok = false
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 500 {
			ok = false
		}
		res = append(res, HTTPCheckResult{
			URL:  url,
			OK:   resp.StatusCode >= 200 && resp.StatusCode < 400,
			Code: resp.StatusCode,
		})
	}
	return
}

func (result HTTPCheckResult) Render(w io.Writer) error {

	return defaultTemplate.Execute(w, result)
}

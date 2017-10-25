package status

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var defaultConfig Config = Config{
	Interval: 10 * time.Second,
}

type CheckResult interface {
	Render(io.Writer) error
}

type Checker interface {
	Check() (bool, []CheckResult, error)
}

type Config struct {
	Interval time.Duration
}

type CheckerWithConfig struct {
	Checker
	Config
}

type Data struct {
	HTMLResults []template.HTML
	Results     []CheckResult
	LastCheck   time.Time
	sync.Mutex
}

type StatusService struct {
	Checkers      []CheckerWithConfig
	IndexTemplate *template.Template
	Data          []Data
}

func New() *StatusService {
	return &StatusService{
		IndexTemplate: defaultIndexTemplate,
	}
}

func (s *StatusService) Add(c Checker) {
	s.Checkers = append(s.Checkers, CheckerWithConfig{c, defaultConfig})
}

func (s *StatusService) AddWith(c Checker, config Config) {
	s.Checkers = append(s.Checkers, CheckerWithConfig{
		Checker: c,
		Config:  config,
	})
}

func (s *StatusService) Run() {
	s.Data = make([]Data, len(s.Checkers))
	for i, checker := range s.Checkers {
		i := i
		checker := checker
		go func() {
			for range time.NewTicker(checker.Interval).C {
				_, res, _ := checker.Check()
				s.Data[i].Results = res

				s.Data[i].HTMLResults = nil
				for _, x := range res {
					buf := &bytes.Buffer{}
					err := x.Render(buf)
					if err != nil {
						panic(err)
					}
					s.Data[i].HTMLResults = append(s.Data[i].HTMLResults, template.HTML(buf.String()))
				}
			}
		}()
	}
}

func (s *StatusService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := s.IndexTemplate.Execute(w, s.Data)
	if err != nil {
		log.Println(err)
	}
}

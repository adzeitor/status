package http_status

import "html/template"

var defaultTemplate = template.Must(template.New("name").Parse(`
	{{if .OK}}
		<li class="list-group-item">
			<div class="d-flex w-100 justify-content-between">
				<h6 class="mb-1 text-truncate">
					{{.URL}}
				</h6>
				<strong>HTTP</strong>
			</div>
			<div class="d-flex">
				<span class="badge badge-secondary">{{.Code}}</span>
			</div>
		</li>
	{{else}}
		<li class="list-group-item list-group-item-danger">
			<div class="d-flex w-100 justify-content-between">
				<h6 class="mb-1 text-truncate">
					{{.URL}}
				</h6>
				<strong>HTTP</strong>
			</div>
			<div class="d-flex">
				<span class="badge badge-secondary">{{.Code}}</span>
				{{.Error}}
			</div>
		</li>
	{{end}}
	`))

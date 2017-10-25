package mysql_status

import "html/template"

var defaultTemplate = template.Must(template.New("name").Parse(`
	{{if .OK}}
		<li class="list-group-item">
			<div class="d-flex w-100 justify-content-between">
				<h6 class="mb-1 text-truncate">
					{{.Name}}
				</h6>
				<strong>MYSQL</strong>
			</div>
		 </li>
	{{else}}
		<li class="list-group-item list-group-item-danger">
			<div class="d-flex w-100 justify-content-between">
				
				<h6 class="mb-1 text-truncate">
					{{.Name}}
				</h6>
				<strong>MYSQL</strong>
			</div>
			<div class="d-flex">
				{{.Error}}
			</div>
		</li>
 	{{end}}
		  
	`))

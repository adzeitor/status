package status

import "html/template"

var defaultIndexTemplate = template.Must(template.New("name").Parse(`

	<html>
	<head>
		<meta http-equiv="refresh" content="2;">
		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css" integrity="sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M" crossorigin="anonymous">
	</head>

	<body>
		<div class="container">
	  		<h1>Status</h1>
	  		<ul class="list-group">
	  		{{ range .}}
				{{ range .HTMLResults }}
					{{.}}
				{{ end }}
			{{end}}
			</ul>
		</div>
	</body>
	</html>
`))

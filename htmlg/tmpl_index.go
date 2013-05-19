package htmlg

import (
	"fmt"
	"html/template"
	"os"
)

type dataIndex struct {
	CountDocuments int
	Documents      []dataIndexDocument
	CountTargets   int
	Targets        []dataIndexTarget
}
type dataIndexDocument struct {
	Name string
	Url  string
}
type dataIndexTarget struct {
	Name string
	Url  string
}

var tmplIndex *template.Template

func init() {
	var err error
	tmplIndex, err = template.New("index").Parse(`
		<div class="row">
			<div class="span4" >
				<h4>{{.CountDocuments}} documents:</h4>
				<ul>
					{{range .Documents}}
						<li><a href="{{.Url}}" >{{.Name}}</a></li>
					{{end}}
				</ul>
			</div>
			<div class="span4" >
				<h4>{{.CountTargets}} targets:</h4>
				<ul>
					{{range .Targets}}
						<li><a href="{{.Url}}" >{{.Name}}</a></li>
					{{end}}
				</ul>
			</div>
		</div>
		`)
	if err != nil {
		fmt.Printf("Error parsing index template: %s\n", err)
		os.Exit(1)
	}
}

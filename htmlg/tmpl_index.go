package htmlg

import (
	"fmt"
	"html/template"
	"os"
)

type dataIndex struct {
	CountDocuments int
	Documents      []dataIndexDocument
}
type dataIndexDocument struct {
	Name string
	Url  string
}

var tmplIndex *template.Template

func init() {
	var err error
	tmplIndex, err = template.New("index").Parse(`
		<p>Threftdoc for {{.CountDocuments}} documents:
			<ul>
				{{range .Documents}}
					<li><a href="{{.Url}}" >{{.Name}}</a></li>
				{{end}}
			</ul>
		</p>
		`)
	if err != nil {
		fmt.Printf("Error parsing index template: %s\n", err)
		os.Exit(1)
	}
}

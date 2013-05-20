package htmlg

import (
	"fmt"
	"html/template"
	"os"
)

var tmplTarget *template.Template

type dataTarget struct {
	Name       string
	Namespaces []dataTargetNamespace
}

type dataTargetNamespace struct {
	Name string
	Url  string
}

func init() {
	var err error
	tmplTarget, err = template.New("target").Parse(`
		<div class="row">
			<div class="span12" >
				<p>The target {{.Name}} has {{len .Namespaces}} namespaces:
				<ul>
					{{range .Namespaces}}
						<li><a href="{{.Url}}" >{{.Name}}</a></li>
					{{end}}
				</ul>
			</div>
		</div>
		`)
	if err != nil {
		fmt.Printf("Error parsing target template: %s\n", err)
		os.Exit(1)
	}
}

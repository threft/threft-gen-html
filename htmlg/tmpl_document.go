package htmlg

import (
	"fmt"
	"html/template"
	"os"
)

var tmplDocument *template.Template

type dataDocument struct {
	Name      string
	Constants []dataDocumentConst
}

type dataDocumentConst struct {
	Name string
	Url  string
}

func init() {
	var err error
	tmplDocument, err = template.New("document").Parse(`
		<div class="row">
			<div class="span4" >
				<h4>Type definitions</h4>
				<ul>
					<li>fdsa</li>
				</ul>
			</div>
			<div class="span4" >
				<h4>Constants</h4>
				<ul>
					{{range .Constants}}
						<li><a href="{{.Url}}" >{{.Name}}</a></li>
					{{end}}
				</ul>
			</div>
			<div class="span4" >
				<h4>Enums</h4>
				<ul>
					<li>fdsa</li>
				</ul>
			</div>
		</div>
		<div class="row">
			<div class="span4" >
				<h4>Structs</h4>
				<ul>
					<li>fdsa</li>
				</ul>
			</div>
			<div class="span4" >
				<h4>Exceptions</h4>
				<ul>
					<li>fdsa</li>
				</ul>
			</div>
			<div class="span4" >
				<h4>Services</h4>
				<ul>
					<li>fdsa</li>
				</ul>
			</div>
		</div>
		`)
	if err != nil {
		fmt.Printf("Error parsing document template: %s\n", err)
		os.Exit(1)
	}
}

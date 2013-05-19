package htmlg

import (
	"fmt"
	"html/template"
	"os"
)

var tmplDocument *template.Template

func init() {
	var err error
	tmplDocument, err = template.New("document").Parse(`
		<div class="row">
			<div class="span4" >
				<h4>Constants</h4>
				<ul>
					<li>fdsa</li>
				</ul>
			</div>
			<div class="span4" >
				<h4>Type definitions</h4>
				<ul>
					<li>fdsa</li>
				</ul>
			</div>
			<div class="span4" >
				<h4>Enums & Senums</h4>
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

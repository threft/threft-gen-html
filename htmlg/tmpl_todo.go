package htmlg

import (
	"fmt"
	"html/template"
	"os"
)

var tmplTodo *template.Template

func init() {
	var err error
	tmplTodo, err = template.New("index").Parse(`
		<div class="container">
			<div class="span12" >
				This page is not being generated yet.<br/>
				<a href="#" onClick="history.go(-1);" class="btn btn-inverse" >Go back</a>
			</div>
		</div>
		`)
	if err != nil {
		fmt.Printf("Error parsing todo template: %s\n", err)
		os.Exit(1)
	}
}

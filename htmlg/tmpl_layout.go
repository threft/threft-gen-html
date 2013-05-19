package htmlg

import (
	"fmt"
	"html/template"
	"os"
)

var tmplHeader *template.Template
var tmplFooter *template.Template

type dataHeader struct {
	Title               string
	PageIndexActive     bool
	PageDocumentsActive bool
	PageTargetsActive   bool
}

func init() {
	var err error
	tmplHeader, err = template.New("header").Parse(`<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>Threftdoc - {{.Title}}</title>
		<meta name="author" content="Generated by threft-gen-html. http://threft.github.io/" >

		<!-- Le HTML5 shim, for IE6-8 support of HTML elements -->
		<!--[if lt IE 9]>
			<script src="http://html5shim.googlecode.com/svn/trunk/html5.js"></script>
		<![endif]-->

		<link href="bootstrap.min.css" rel="stylesheet" type="text/css" media="all"  >
		<link href="style.css" rel="stylesheet" type="text/css" media="all" >
	</head>

	<body>
		<div class="navbar navbar-inverse navbar-fixed-top">
			<div class="navbar-inner">
				<div class="container">
					<a class="brand" href="./index.html">Threftdoc</a>
					<ul class="nav">
						<li {{if .PageIndexActive}} class="active"{{end}}><a href="./index.html" title="Index" >Index</a></li>
						<li {{if .PageDocumentsActive}} class="active"{{end}}><a href="./documents.html" title="Documents" >Documents</a></li>
						<li {{if .PageTargetsActive}} class="active"{{end}}><a href="./targets.html" title="Targets" >Targets</a></li>
					</ul>
				</div>
			</div>
		</div>
		<div class="container">
			`)
	if err != nil {
		fmt.Printf("Error parsing header template: %s\n", err)
		os.Exit(1)
	}

	tmplFooter, err = template.New("footer").Parse(`
		</div>
		<div class="container">
			<hr>
			<footer>
				<p>
					this threftdoc was generated by threft-gen-html &bull; <a href="http://threft.github.io" >threft.github.io</a> &bull; built with <a href="http://twitter.github.com/bootstrap/" target="_blank">twitter bootstrap</a>
				</p>
			</footer>
		</div>
	</body>
</html>`)
	if err != nil {
		fmt.Printf("Error parsing footer template: %s\n", err)
		os.Exit(1)
	}
}

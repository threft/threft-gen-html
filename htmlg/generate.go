package htmlg

import (
	"fmt"
	"github.com/threft/threft/tidm"
	"html/template"
	"os"
	"regexp"
	"sort"
)

// entry point for gog
func GenerateHtml(t *tidm.TIDM) {
	err := writeStaticFiles()
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	generateIndexPage(t)

	generateDocumentPage(t)
	generateDocumentPages(t)

	generateTargetPage(t)
	generateTargetPages(t)

	generateDefinitionConstPages(t)

	fmt.Println("Thank you for using threft-gen-html.")
}

func writePage(fileName string, thd *dataHeader, contentTemplate *template.Template, contentData interface{}) {
	pageFile, err := os.Create(fileName + ".html")
	if err != nil {
		fmt.Printf("Error creating file '%s'. %s\n", fileName, err)
	}

	err = tmplHeader.Execute(pageFile, thd)
	if err != nil {
		fmt.Printf("Error executing tmplHeader. %s\n", err)
		os.Exit(1)
	}

	err = contentTemplate.Execute(pageFile, contentData)
	if err != nil {
		fmt.Printf("Error executing contentTemplate. %s\n", err)
		os.Exit(1)
	}

	err = tmplFooter.Execute(pageFile, nil)
	if err != nil {
		fmt.Printf("Error executing tmplFooter. %s\n", err)
		os.Exit(1)
	}
}

// generates targets.html
func generateTargetPage(t *tidm.TIDM) {
	writePage("targets", &dataHeader{Title: "Targets (TODO)"}, tmplTodo, nil)
}

// generates target-tName.html
func generateTargetPages(t *tidm.TIDM) {
	for targetName, _ := range t.Targets {
		if targetName == "*" {
			targetName = "* (default)"
		}
		dataHeaderTodo := &dataHeader{
			Title: string(targetName) + " (TODO)",
		}
		writePage("target-"+urlify(string(targetName)), dataHeaderTodo, tmplTodo, nil)
	}
}

// generates documents.html
func generateDocumentPage(t *tidm.TIDM) {
	writePage("documents", &dataHeader{Title: "Documents (TODO)"}, tmplTodo, nil)
}

// generates document-dName.html
func generateDocumentPages(t *tidm.TIDM) {
	for docName, doc := range t.Documents {
		dataHeader := &dataHeader{
			Title: "Document - " + string(docName),
		}
		dataDocument := &dataDocument{
			Name: string(docName),
		}

		// prepare constant data
		constNames := []string{}
		for identifierName, _ := range doc.Consts {
			constNames = append(constNames, string(identifierName))
		}
		sort.Strings(constNames)
		for _, constName := range constNames {
			dataDocument.Constants = append(dataDocument.Constants, dataDocumentConst{
				Name: constName,
				Url:  fmt.Sprintf("definition-const-%s-%s.html", urlify(string(docName)), urlify(constName)),
			})
		}

		// write document page
		writePage("document-"+urlify(string(docName)), dataHeader, tmplDocument, dataDocument)
	}
}

// generates definition-const-docName-identifierName.html
func generateDefinitionConstPages(t *tidm.TIDM) {
	for docName, doc := range t.Documents {
		for identifierName, _ := range doc.Consts {
			dataHeader := &dataHeader{
				Title: fmt.Sprintf("Constant definition - %s - (%s)", identifierName, docName),
			}
			writePage(fmt.Sprintf("definition-const-%s-%s", urlify(string(docName)), urlify(string(identifierName))), dataHeader, tmplTodo, nil)
		}
	}
}

// generates index.html
func generateIndexPage(t *tidm.TIDM) {
	dataHeader := &dataHeader{
		Title: "index",
	}

	data := &dataIndex{
		CountDocuments: len(t.Documents),
		CountTargets:   len(t.Targets),
	}

	// prepare data on documents
	docNames := []string{}
	for docName, _ := range t.Documents {
		docNames = append(docNames, string(docName))
	}
	sort.Strings(docNames)
	for _, docName := range docNames {
		data.Documents = append(data.Documents, dataIndexDocument{
			Name: docName,
			Url:  "document-" + urlify(docName) + ".html",
		})
	}

	// prepare data on targets
	targetNames := []string{}
	for targetName, _ := range t.Targets {
		if targetName == "*" {
			targetName = "* (default)"
		}
		targetNames = append(targetNames, string(targetName))
	}
	sort.Strings(targetNames)
	for _, targetName := range targetNames {
		data.Targets = append(data.Targets, dataIndexTarget{
			Name: targetName,
			Url:  fmt.Sprintf("target-%s.html", urlify(targetName)),
		})
	}

	writePage("index", dataHeader, tmplIndex, data)
}

// characters to replace with an underscore to have pretty url's
var regexpUrlify = regexp.MustCompile(`[\. ]`)

// replace certain characters with underscore to have a pretty url
func urlify(input string) string {
	return regexpUrlify.ReplaceAllString(input, "_")
}

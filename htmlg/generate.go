package htmlg

import (
	"fmt"
	"github.com/threft/threft/tidm"
	"html/template"
	"os"
	"regexp"
	"sort"
)

var (
	dataPagesDocuments []dataPageDocument
	dataPagesTargets   []dataPageTarget
)

type dataPageDocument struct {
	Name string
	Url  string
}
type dataPageTarget struct {
	Name string
	Url  string
}

// entry point for gog
func GenerateHtml(t *tidm.TIDM) {
	err := writeStaticFiles()
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	preparePagesData(t)

	generateIndexPage(t)

	generateDocumentPages(t)

	generateTargetPages(t)

	generateDefinitionConstPages(t)

	fmt.Println("Thank you for using threft-gen-html.")
}

func preparePagesData(t *tidm.TIDM) {
	// prepare data on documents
	docNames := []string{}
	for docName, _ := range t.Documents {
		docNames = append(docNames, string(docName))
	}
	sort.Strings(docNames)
	for _, docName := range docNames {
		dataPagesDocuments = append(dataPagesDocuments, dataPageDocument{
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
		dataPagesTargets = append(dataPagesTargets, dataPageTarget{
			Name: targetName,
			Url:  fmt.Sprintf("target-%s.html", urlify(targetName)),
		})
	}
}

func writePage(fileName string, title string, contentTemplate *template.Template, contentData interface{}) {
	// prepare dataHeader object
	dataHeader := &dataHeader{
		Title:     title,
		Documents: dataPagesDocuments,
		Targets:   dataPagesTargets,
	}

	pageFile, err := os.Create(fileName + ".html")
	if err != nil {
		fmt.Printf("Error creating file '%s'. %s\n", fileName, err)
	}

	err = tmplHeader.Execute(pageFile, dataHeader)
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

// generates target-tName.html
func generateTargetPages(t *tidm.TIDM) {
	for targetName, _ := range t.Targets {
		if targetName == "*" {
			targetName = "* (default)"
		}
		pageTitle := string(targetName) + " (TODO)"
		writePage("target-"+urlify(string(targetName)), pageTitle, tmplTodo, nil)
	}
}

// generates document-dName.html
func generateDocumentPages(t *tidm.TIDM) {
	for docName, doc := range t.Documents {
		pageTitle := "Document - " + string(docName)
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
		writePage("document-"+urlify(string(docName)), pageTitle, tmplDocument, dataDocument)
	}
}

// generates definition-const-docName-identifierName.html
func generateDefinitionConstPages(t *tidm.TIDM) {
	for docName, doc := range t.Documents {
		for identifierName, _ := range doc.Consts {
			pageTitle := fmt.Sprintf("Constant definition - %s - (%s)", identifierName, docName)
			writePage(fmt.Sprintf("definition-const-%s-%s", urlify(string(docName)), urlify(string(identifierName))), pageTitle, tmplTodo, nil)
		}
	}
}

// generates index.html
func generateIndexPage(t *tidm.TIDM) {
	data := &dataIndex{
		CountDocuments: len(t.Documents),
		Documents:      dataPagesDocuments,
		CountTargets:   len(t.Targets),
		Targets:        dataPagesTargets,
	}

	writePage("index", "index", tmplIndex, data)
}

// characters to replace with an underscore to have pretty url's
var regexpUrlify = regexp.MustCompile(`[\. ]`)

// replace certain characters with underscore to have a pretty url
func urlify(input string) string {
	return regexpUrlify.ReplaceAllString(input, "_")
}

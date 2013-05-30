package htmlg

import (
	"fmt"
	"github.com/threft/threft/tidm"
	"html/template"
	"os"
	"regexp"
	"sort"
)

var generatedPages = 0

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

	generateNamespacePages(t)

	generateDefinitionTypedefPages(t)

	generateDefinitionConstPages(t)

	generateDefinitionEnumPages(t)

	generateDefinitionStructPages(t)

	generateDefinitionServicePages(t)

	fmt.Printf("Generated %d html pages.\n", generatedPages)
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

func writePage(fileName string, title string, pateTemplate *template.Template, pageData interface{}) {
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

	err = pateTemplate.Execute(pageFile, pageData)
	if err != nil {
		fmt.Printf("Error executing page template. %s\n", err)
		os.Exit(1)
	}

	err = tmplFooter.Execute(pageFile, nil)
	if err != nil {
		fmt.Printf("Error executing tmplFooter. %s\n", err)
		os.Exit(1)
	}

	generatedPages++
}

// generates target-tName.html
func generateTargetPages(t *tidm.TIDM) {
	for targetName, target := range t.Targets {
		if targetName == "*" {
			targetName = "* (default)"
		}
		pageTitle := string(targetName)
		dataTarget := &dataTarget{
			Name: string(targetName),
		}
		namespaceNames := []string{}
		for namespaceName, _ := range target.Namespaces {
			namespaceNames = append(namespaceNames, string(namespaceName))
		}
		sort.Strings(namespaceNames)
		for _, namespaceName := range namespaceNames {
			dataTarget.Namespaces = append(dataTarget.Namespaces, dataTargetNamespace{
				Name: namespaceName,
				Url:  fmt.Sprintf("namespace-%s-%s.html", urlify(string(targetName)), urlify(namespaceName)),
			})
		}
		writePage("target-"+urlify(string(targetName)), pageTitle, tmplTarget, dataTarget)
	}
}

// generates namespace-tName-nName.html
func generateNamespacePages(t *tidm.TIDM) {
	for targetName, target := range t.Targets {
		if targetName == "*" {
			targetName = "* (default)"
		}
		for namespaceName, _ := range target.Namespaces {
			pageTitle := string(targetName) + " - " + string(namespaceName)
			writePage("namespace-"+urlify(string(targetName))+"-"+urlify(string(namespaceName)), pageTitle, tmplTodo, nil)
		}
	}
}

// generates document-dName.html
func generateDocumentPages(t *tidm.TIDM) {
	for docName, doc := range t.Documents {
		pageTitle := "Document - " + string(docName)
		dataDocument := &dataDocument{
			Name: string(docName),
		}

		// prepare typedef data
		typedefNames := []string{}
		for identifierName, _ := range doc.Typedefs {
			typedefNames = append(typedefNames, string(identifierName))
		}
		sort.Strings(typedefNames)
		for _, typedefName := range typedefNames {
			dataDocument.Typedefs = append(dataDocument.Typedefs, dataDocumentTypedef{
				Name: typedefName,
				Url:  fmt.Sprintf("definition-typedef-%s-%s.html", urlify(string(docName)), urlify(typedefName)),
			})
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

// generates definition-typedef-docName-identifierName.html
func generateDefinitionTypedefPages(t *tidm.TIDM) {
	for docName, doc := range t.Documents {
		for identifierName, _ := range doc.Typedefs {
			pageTitle := fmt.Sprintf("Typedef definition - %s - (%s)", identifierName, docName)
			writePage(fmt.Sprintf("definition-typedef-%s-%s", urlify(string(docName)), urlify(string(identifierName))), pageTitle, tmplTodo, nil)
		}
	}
}

// generates definition-const-docName-identifierName.html
func generateDefinitionConstPages(t *tidm.TIDM) {
	for docName, doc := range t.Documents {
		for identifierName, _ := range doc.Consts {
			pageTitle := fmt.Sprintf("Const definition - %s - (%s)", identifierName, docName)
			writePage(fmt.Sprintf("definition-const-%s-%s", urlify(string(docName)), urlify(string(identifierName))), pageTitle, tmplTodo, nil)
		}
	}
}

// generates definition-enum-docName-identifierName.html
func generateDefinitionEnumPages(t *tidm.TIDM) {
	for docName, doc := range t.Documents {
		for identifierName, _ := range doc.Enums {
			pageTitle := fmt.Sprintf("Enum definition - %s - (%s)", identifierName, docName)
			writePage(fmt.Sprintf("definition-enum-%s-%s", urlify(string(docName)), urlify(string(identifierName))), pageTitle, tmplTodo, nil)
		}
	}
}

// generates definition-struct-docName-identifierName.html
func generateDefinitionStructPages(t *tidm.TIDM) {
	for docName, doc := range t.Documents {
		for identifierName, _ := range doc.Structs {
			pageTitle := fmt.Sprintf("Struct definition - %s - (%s)", identifierName, docName)
			writePage(fmt.Sprintf("definition-struct-%s-%s", urlify(string(docName)), urlify(string(identifierName))), pageTitle, tmplTodo, nil)
		}
	}
}

// generates definition-service-docName-identifierName.html
func generateDefinitionServicePages(t *tidm.TIDM) {
	for docName, doc := range t.Documents {
		for identifierName, _ := range doc.Services {
			pageTitle := fmt.Sprintf("Service definition - %s - (%s)", identifierName, docName)
			writePage(fmt.Sprintf("definition-service-%s-%s", urlify(string(docName)), urlify(string(identifierName))), pageTitle, tmplTodo, nil)
		}
	}
}

// generates index.html
func generateIndexPage(t *tidm.TIDM) {
	data := &dataIndex{
		Documents: dataPagesDocuments,
		Targets:   dataPagesTargets,
	}

	writePage("index", "index", tmplIndex, data)
}

// characters to replace with an underscore to have pretty url's
var regexpUrlify = regexp.MustCompile(`[\. ]`)

// replace certain characters with underscore to have a pretty url
func urlify(input string) string {
	return regexpUrlify.ReplaceAllString(input, "_")
}

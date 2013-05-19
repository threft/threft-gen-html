package htmlg

import (
	"fmt"
	"github.com/threft/threft/tidm"
	"html/template"
	"os"
	"regexp"
	"sort"
)

// urls:
// index
// documents
// document-dname
// targets
// target-tname
// namespace-tname-nname
// definition-dname-identifiername

func GenerateHtml(t *tidm.TIDM) {
	err := writeStaticFiles()
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	err = writeIndex(t)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

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

func writeIndex(t *tidm.TIDM) error {
	dataHeader := &dataHeader{
		Title: "index",
	}

	data := &dataIndex{
		CountDocuments: len(t.Documents),
	}
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

	writePage("index", dataHeader, tmplIndex, data)

	return nil
}

var regexpUrlify = regexp.MustCompile(`[\.]`)

func urlify(input string) string {
	return regexpUrlify.ReplaceAllString(input, "_")
}

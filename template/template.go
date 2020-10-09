package template

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"text/template"

	"../config"
	"../search-condition"
)

type IndexTemplateValues struct {
	ModifyDate         string
	SearchResults      []search.SearchResult
	SearchResultsCount int
}

func NewIndexTempalteValues() *IndexTemplateValues {
	return &IndexTemplateValues{}
}

func (templateValues *IndexTemplateValues) MakeIndexTemplate() (string, error) {
	buffer := new(bytes.Buffer)

	temp := template.Must(template.ParseFiles(config.C.Files.Template_File))
	if err := temp.Execute(buffer, templateValues); err != nil {
		return "", err
	}

	filename := fmt.Sprintf(config.C.Files.Output_File, templateValues.ModifyDate)
	if err := ioutil.WriteFile(filename, buffer.Bytes(), 0666); err != nil {
		return "", err
	}
	return filename, nil
}

package template

import (
	"bytes"
	"io/ioutil"
	"text/template"

	"../youtube"
)

type IndexTemplateValues struct {
	Article    *youtube.YoutubeArticle
	ModifyDate string
}

func NewIndexTempalteValues() *IndexTemplateValues {
	return &IndexTemplateValues{}
}

func (templateValues *IndexTemplateValues) MakeIndexTemplate() (string, error) {
	buffer := new(bytes.Buffer)

	temp := template.Must(template.ParseFiles("template/template.html"))
	if err := temp.Execute(buffer, templateValues); err != nil {
		return "", err
	}

	filename := "index-" + templateValues.ModifyDate + ".html"
	if err := ioutil.WriteFile(filename, buffer.Bytes(), 0666); err != nil {
		return "", err
	}
	return filename, nil
}

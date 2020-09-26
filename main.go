package main

import (
	"bufio"
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"time"

	"./youtube"
)

type TemplateValues struct {
	Article    *youtube.YoutubeArticle
	ModifyDate string
}

func main() {
	productKey, err := getProductKey()
	if err != nil {
		log.Fatal(err)
	}

	today := time.Now()
	publishedAfter := today.AddDate(0, 0, -7).UTC().Format(time.RFC3339)
	modifyDate := today.Format("2006-01-02")

	youtubeApi := youtube.NewYoutube()
	article, err := youtubeApi.SetProductKey(productKey).SetMaxResults(50).
		SetPublishedAfter(publishedAfter).Connect().GetArticle()
	if err != nil {
		log.Fatal(err)
	}

	templateValues := TemplateValues{}
	templateValues.Article = article
	templateValues.ModifyDate = modifyDate

	buffer := new(bytes.Buffer)

	t := template.Must(template.ParseFiles("template/template.html"))
	if err := t.Execute(buffer, templateValues); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("index.html", buffer.Bytes(), 0666); err != nil {
		log.Fatal(err)
	}
}

func getProductKey() (string, error) {
	fp, err := os.Open("./product_key.conf")
	if err != nil {
		return "", err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	return scanner.Text(), nil
}

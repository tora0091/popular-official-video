package main

import (
	"bufio"
	"bytes"
	"html/template"
	"io"
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

	filename := "index-" + templateValues.ModifyDate + ".html"
	if err := ioutil.WriteFile(filename, buffer.Bytes(), 0666); err != nil {
		log.Fatal(err)
	}

	if err = copyIndexHtml(filename); err != nil {
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

func copyIndexHtml(filename string) error {
	src := filename
	dst := "../../firebase/public/index.html"

	fsrc, err := os.OpenFile(src, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer fsrc.Close()

	fdst, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer fdst.Close()

	_, err = io.Copy(fdst, fsrc)
	if err != nil {
		return err
	}
	return nil
}

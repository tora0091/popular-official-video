package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"

	"./template"
	"./youtube"
)

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

	templateValues := template.NewIndexTempalteValues()
	templateValues.Article = article
	templateValues.ModifyDate = modifyDate
	filename, err := templateValues.MakeIndexTemplate()
	if err != nil {
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

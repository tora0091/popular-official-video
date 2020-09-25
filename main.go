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

func main() {
	productKey, err := getProductKey()
	if err != nil {
		log.Fatal(err)
	}

	publishedAfter := time.Now().AddDate(0, 0, -7).UTC().Format(time.RFC3339)

	youtubeApi := youtube.NewYoutube()
	article, err := youtubeApi.SetProductKey(productKey).SetMaxResults(50).
		SetPublishedAfter(publishedAfter).Connect().GetArticle()
	if err != nil {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)

	t := template.Must(template.ParseFiles("template/template.html"))
	if err := t.Execute(buffer, article); err != nil {
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

package main

import (
	"bufio"
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"./youtube"
)

func main() {
	productKey, err := getProductKey()
	if err != nil {
		log.Fatal(err)
	}

	youtubeApi := youtube.NewYoutube()
	youtubeApi.ProductKey = productKey
	youtubeApi.MaxResults = 50
	article, err := youtubeApi.Connect().GetArticle()
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

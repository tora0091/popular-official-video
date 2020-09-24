package main

import (
	"bufio"
	"fmt"
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
	article, err := youtubeApi.Connect().GetArticle()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(article)
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

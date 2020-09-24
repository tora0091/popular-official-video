package main

import (
	"fmt"
	"log"

	"./youtube"
)

func main() {
	youtubeApi := youtube.NewYoutube()
	article, err := youtubeApi.Connect().GetArticle()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(article)
}

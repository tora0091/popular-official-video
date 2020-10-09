package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"

	"./config"
	"./search-condition"
	"./template"
)

func init() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	productKey, err := getProductKey()
	if err != nil {
		log.Fatal(err)
	}

	today := time.Now()
	publishedAfter := today.AddDate(0, 0, -7).UTC().Format(time.RFC3339)
	modifyDate := today.Format("2006-01-02")

	searchConditions := search.NewSearchCondition()
	searchResults := search.NewSearchResult()
	for _, condition := range searchConditions {
		article, err := condition.GetSearchContents(productKey, publishedAfter)
		if err != nil {
			log.Print(err)
		}

		searchResults = append(searchResults,
			search.SearchResult{
				Code:    condition.Code,
				Word:    condition.Word,
				Article: article,
			})
	}

	templateValues := template.NewIndexTempalteValues()
	templateValues.ModifyDate = modifyDate
	templateValues.SearchResults = searchResults
	templateValues.SearchResultsCount = len(searchResults)

	filename, err := templateValues.MakeIndexTemplate()
	if err != nil {
		log.Fatal(err)
	}

	if err = copyIndexHtml(filename); err != nil {
		log.Fatal(err)
	}
}

func getProductKey() (string, error) {
	fp, err := os.Open(config.C.Files.Product_Conf)
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
	dst := config.C.Files.Index_Html

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

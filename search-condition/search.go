package search

import (
	"../youtube"
)

type SearchCondition struct {
	Code string
	Word string
}

type SearchResult struct {
	Code    string
	Word    string
	Article *youtube.YoutubeArticle
}

func NewSearchCondition() []SearchCondition {
	return []SearchCondition{
		{
			Code: "GSW",
			Word: "Golden State Warriors",
		},
		{
			Code: "PHI",
			Word: "Philadelphia 76ers",
		},
	}
}

func NewSearchResult() []SearchResult {
	return []SearchResult{}
}

func (condition SearchCondition) GetSearchContents(productKey, publishedAfter string) (*youtube.YoutubeArticle, error) {
	youtubeApi := youtube.NewYoutube()
	article, err := youtubeApi.SetProductKey(productKey).SetSearchWord(condition.Word + " Rumors").
		SetPublishedAfter(publishedAfter).SetOrder("relevance").Connect().GetArticle()
	if err != nil {
		return nil, err
	}
	return article, nil
}

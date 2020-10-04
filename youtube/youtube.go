package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	BASE_URL        = "https://www.googleapis.com/youtube/v3/search"
	PRODUCT_KEY     = ""
	PART            = "id,snippet"
	MAX_RESULTS     = 10
	ORDER           = "viewCount"
	Q               = "official+video"
	REGION_CODE     = "us"
	TYPE            = "video"
	PULBISHED_AFTER = ""
)

type YoutubeApi struct {
	BaseUrl        string
	ProductKey     string
	Part           string
	MaxResults     int
	Order          string
	Q              string
	RegionCode     string
	Type           string
	PublishedAfter string
	ApiUrl         string
}

func NewYoutube() *YoutubeApi {
	return &YoutubeApi{
		BaseUrl:        BASE_URL,
		ProductKey:     PRODUCT_KEY,
		Part:           PART,
		MaxResults:     MAX_RESULTS,
		Order:          ORDER,
		Q:              Q,
		RegionCode:     REGION_CODE,
		Type:           TYPE,
		PublishedAfter: PULBISHED_AFTER,
	}
}

type YoutubeArticle struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	RegionCode    string `json:"regionCode"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind string `json:"kind"`
		Etag string `json:"etag"`
		ID   struct {
			Kind    string `json:"kind"`
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
			} `json:"thumbnails"`
			ChannelTitle         string    `json:"channelTitle"`
			LiveBroadcastContent string    `json:"liveBroadcastContent"`
			PublishTime          time.Time `json:"publishTime"`
		} `json:"snippet"`
	} `json:"items"`
}

func (y *YoutubeApi) GetArticle() (*YoutubeArticle, error) {
	resp, err := http.Get(y.ApiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Status Code is %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var article YoutubeArticle
	if err := json.Unmarshal(body, &article); err != nil {
		return nil, err
	}
	return &article, nil
}

func (y *YoutubeApi) Connect() *YoutubeApi {
	maxResults := strconv.Itoa(y.MaxResults)
	url := y.BaseUrl + "?part=" + y.Part + "&maxResults=" + maxResults +
		"&order=" + y.Order + "&q=" + y.Q + "&regionCode=" + y.RegionCode + "&type=" + y.Type + "&key=" + y.ProductKey
	if y.PublishedAfter != "" {
		url = url + "&publishedAfter=" + y.PublishedAfter
	}
	y.ApiUrl = url
	return y
}

func (y *YoutubeApi) SetProductKey(pkey string) *YoutubeApi {
	y.ProductKey = pkey
	return y
}

func (y *YoutubeApi) SetMaxResults(maxResults int) *YoutubeApi {
	y.MaxResults = maxResults
	return y
}

func (y *YoutubeApi) SetRegionCode(regionCode string) *YoutubeApi {
	y.RegionCode = regionCode
	return y
}

func (y *YoutubeApi) SetPublishedAfter(publishedAfter string) *YoutubeApi {
	y.PublishedAfter = publishedAfter
	return y
}

func (y *YoutubeApi) SetSearchWord(searchWord string) *YoutubeApi {
	y.Q = url.QueryEscape(searchWord)
	return y
}

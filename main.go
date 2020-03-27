package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	api := anaconda.NewTwitterApiWithCredentials(
		os.Getenv("TWITTER_ACCESS_TOKEN"),
		os.Getenv("TWITTER_ACCESS_SECRET"),
		os.Getenv("TWITTER_CONSUMER_KEY"),
		os.Getenv("TWITTER_SECRET_KEY"),
	)

	v := url.Values{}
	v.Set("count", "10")
	v.Set("lang", "pt")

	searchResult, _ := api.GetSearch("#golang -filter:retweets", v)
	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text, tweet.CreatedAt, tweet.Id)
	}
}

package tw

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

// Client is twitter client
type Client struct {
	api    *anaconda.TwitterApi
	logger *log.Logger
}

// NewClient returns api client object
func NewClient() *Client {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	logger := log.New(os.Stdout, "logger: ", log.Lshortfile)
	c := new(Client)
	c.api = api
	c.logger = logger
	return c
}

// Tweet tweets arg
func (c *Client) Tweet(content string) {
	_, err := c.api.PostTweet(os.Args[1], nil)
	if err != nil {
		panic(err)
	}
}

// Stream shows timeline
func (c *Client) Stream() {
	c.logger.Println("Stream start")
	v := url.Values{}
	stream := c.api.UserStream(v)

	for {
		select {
		case item := <-stream.C:
			switch status := item.(type) {
			case anaconda.Tweet:
				fmt.Printf("@%s at %s\n", status.User.ScreenName, getTime())
				fmt.Printf("%s\n", status.Text)
				fmt.Printf("\n")
			}
		}
	}
}

func getTime() string {
	const f = "2006-01-02 15:04:05"
	t := time.Now()
	return t.Format(f)
}

// Usage returns usage string
func Usage() string {
	return `Usage:
  tw				show your time line as stream
  tw "your_tweet_content"	post tweet
	`
}

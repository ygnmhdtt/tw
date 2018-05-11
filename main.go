package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

// client is twitter client
type client struct {
	api    *anaconda.TwitterApi
	logger *log.Logger
}

// newClient returns api client object
func newClient() *client {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	logger := log.New(os.Stdout, "logger: ", log.Lshortfile)
	c := new(client)
	c.api = api
	c.logger = logger
	return c
}

func main() {
	c := newClient()
	if 3 <= len(os.Args) {
		fmt.Println(usage())
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		c.stream()
	} else {
		c.tweet(os.Args[1])
	}
}

// tweet tweets arg
func (c *client) tweet(content string) {
	_, err := c.api.PostTweet(os.Args[1], nil)
	if err != nil {
		panic(err)
	}
}

// stream shows timeline
func (c *client) stream() {
	c.logger.Println("Stream start")
	v := url.Values{}
	stream := c.api.UserStream(v)

	for {
		select {
		case item := <-stream.C:
			switch status := item.(type) {
			case anaconda.Tweet:
				fmt.Printf("@%s (%s) at %s\n", status.User.Name, status.User.ScreenName, getTime())
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

// usage returns usage string
func usage() string {
	return `Usage:
  tw				show your time line as stream
  tw "your_tweet_content"	post tweet
	`
}

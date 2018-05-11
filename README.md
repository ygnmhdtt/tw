tw - minimum twitter client for commandline

### features

* post tweet
* show timeline streaming

### authentication

You need to create twitter app [here](https://apps.twitter.com).
And, you must prepare 4 environment variables:

```
export TWITTER_CONSUMER_KEY="AAAAAAA"
export TWITTER_CONSUMER_SECRET="AAAAAAA"
export TWITTER_ACCESS_TOKEN="AAAAAAA"
export TWITTER_ACCESS_TOKEN_SECRET="AAAAAAA"
```

### Installation

```
$ go get github.com/ygnmhdtt/tw/cmd/tw
```

### Usage

* Streaming your timeline

```
$ tw
```

* Post new tweet

```
$ tw "your awesome tweet!!"
```

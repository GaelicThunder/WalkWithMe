package twitterconnector

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type Twitter interface {
	TweetStatusUpdate(tweet string) error
}

func NewTwitter() Twitter {
	if os.Getenv("mode_mock") != "" {
		log.Printf("The walker will work in mock mode")
		return &MockedTwitter{}
	}
	return &ConcreteTwitter{}
}

type ConcreteTwitter struct{}

// tweetStatusUpdate update on twitter the given message
func (t *ConcreteTwitter) TweetStatusUpdate(tweet string) error {
	clientID := os.Getenv("TWITTER_ID")
	secret := os.Getenv("TWITTER_SECRET")
	accessToken := os.Getenv("ACCESS_TKN")
	accessSecret := os.Getenv("ACCESS_SECRET")
	
	config := oauth1.NewConfig(clientID, secret)
	token := oauth1.NewToken(accessToken, accessSecret)
	
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	_, _, err := client.Statuses.Update(tweet, nil)
	return err
}

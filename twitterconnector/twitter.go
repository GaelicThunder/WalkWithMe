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
	clientID := os.Getenv("TWITER_ID")
	secret := os.Getenv("TWITER_SECRET")
	tknURL := os.Getenv("TWITER_URL")
	// oauth2 configures a client that uses app credentials to keep a fresh token
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: secret,
		TokenURL:     tknURL,
	}
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	// Twitter client
	client := twitter.NewClient(httpClient)

	_, _, err := client.Statuses.Update(tweet, nil)
	return err
}

package twitterconnector

import "log"

type MockedTwitter struct{}

// tweetStatusUpdate update on twitter the given message
func (m *MockedTwitter) TweetStatusUpdate(tweet string) error {
	log.Print(tweet)
	return nil
}

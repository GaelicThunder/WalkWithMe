package main

import (
	"fmt"
	"log"
	"time"
	"walkwithme/database"
	"walkwithme/twitterconnector"
	"walkwithme/walker"
)

func main() {
	// initialise database and get the actual walk status
	db := database.NewDynamoDB()
	walk, err := db.GetWalk()
	if err != nil {
		log.Fatalf("Error while get walk: %s\n", err.Error())
	}
	// define if we should walk or rest
	if time.Since(walk.LastRest) > 20*time.Hour && !walk.LastRest.IsZero() {
		err = twitterconnector.NewTwitter().TweetStatusUpdate("I'm going to rest! Cheers!")
		if err != nil {
			log.Fatalf("fatal error: %s", err)
		}
		walk.LastRest = time.Now()
		db.SaveWalk(walk)
		return
	}
	if time.Since(walk.LastRest) < 8*time.Hour {
		// we are resting! No interuption allowed!
		return
	}
	// we have enought energy to keep going!
	actualPosition, timeWalked, err := walker.NewWalker().PositionComputation(walk.ActualPosition, walk.To)
	if err != nil {
		log.Fatalf("Error while get walk: %s\n", err.Error())
	}
	walk.ActualPosition = actualPosition
	walk.TotalHoursWalked += timeWalked
	// tweet status update
	err = twitterconnector.NewTwitter().TweetStatusUpdate(fmt.Sprintf("Starting from %s after %s time I'm now at %s", walk.From, walk.TotalHoursWalked, walk.ActualPosition))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	// since we have done everything we can store the updated status to the database
	db.SaveWalk(walk)
}

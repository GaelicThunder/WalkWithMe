package main

import (
	"fmt"
	"log"
	"time"
	"walkwithme/database"
	"walkwithme/twitterconnector"
	"walkwithme/walker"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(processing)
}

func processing(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
			log.Fatalf("error while twitting rest: %s", err)
		}
		walk.LastRest = time.Now()
		db.SaveWalk(walk)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 200}, nil
	}
	if time.Since(walk.LastRest) < 8*time.Hour {
		// we are resting! No interuption allowed!
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 200}, nil
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
		log.Fatalf("error while twitting: %s", err)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 200}, nil
	}

	// since we have done everything we can store the updated status to the database
	err = db.SaveWalk(walk)
	if err != nil {
		log.Fatalf("error while storing walk status: %s", err)
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 200}, nil
	}
	return events.APIGatewayProxyResponse{Body: "", StatusCode: 200}, nil
}

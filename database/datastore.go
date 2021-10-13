package database

import (
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// WalkStatus is the struct stored on
// the database which will contain all the info about the walks,
// the one with status "actual" is the walk we have to updated
type WalkStatus struct {
	ID               string
	From             string
	To               string
	ActualPosition   string
	Status           string
	TotalHoursWalked time.Duration
	LastRest         time.Time
}

// Database rappresent an abstraction to the database
type Database interface {
	GetWalk() (*WalkStatus, error)
	SaveWalk(*WalkStatus) error
}

// DyanamoDB is responsible to communicate to dynamoDB
type DyanamoDB struct {
	db *dynamo.DB
}

// GetWalk return the actual walk in progress
func (d *DyanamoDB) GetWalk() (*WalkStatus, error) {
	var walk WalkStatus
	err := d.db.Table("walkwithme").Get("Status", "actual").One(&walk)
	if err != nil {
		return nil, err
	}
	return &walk, nil
}

// SaveWalk save the actual walk status
func (d *DyanamoDB) SaveWalk(walk *WalkStatus) error {
	err := d.db.Table("gotchi").Put(walk).Run()
	if err != nil {
		return err
	}
	return nil
}

// NewDynamoDB return a dynamoDB object which respect the Database interface
func NewDynamoDB() Database {
	if os.Getenv("mode_mock") != "" {
		log.Printf("The database will work in mock mode")
		return NewMockDatabase()
	}
	session, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	})
	if err != nil {
		log.Printf("Error while setup the session: %s\n", err.Error())
		os.Exit(1)
	}
	db := dynamo.New(session)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Printf("Session to dynamo establish.\n")

	return &DyanamoDB{
		db: db,
	}
}

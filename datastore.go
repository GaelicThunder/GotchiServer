package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// GotchiInformation is the struct stored on
// the database which will contain all the info about gotchi
type GotchiInformation struct {
	ID         string
	KnowGotchi []string
}

// Database rappresent an abstraction to the database
type Database interface {
	GetGotchi(string) ([]string, error)
	SaveGotchi(string, string) error
}

// DyanamoDB is responsible to communicate to dynamoDB
type DyanamoDB struct {
	db *dynamo.DB
}

// GetGotchi return the list of know gotchi
func (d *DyanamoDB) GetGotchi(myGotchiID string) ([]string, error) {
	var myself GotchiInformation
	err := d.db.Table("gotchi").Get("ID", myGotchiID).One(&myself)
	if err != nil {
		return nil, err
	}
	return myself.KnowGotchi, nil
}

// SaveGotchi save a new know gotchiID in the database
func (d *DyanamoDB) SaveGotchi(myGotchiID, newGotchiID string) error {
	var myself GotchiInformation
	err := d.db.Table("gotchi").Get("ID", myGotchiID).One(&myself)
	if err != nil {
		return err
	}
	myself.KnowGotchi = append(myself.KnowGotchi, newGotchiID)
	err = d.db.Table("gotchi").Put(myself).Run()
	if err != nil {
		return err
	}
	return nil
}

// NewDynamoDB return a dynamoDB object which respect the Database interface
func NewDynamoDB() Database {
	awsID := os.Getenv("AWS_ACCESS_KEY_ID")
	awsKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: credentials.NewStaticCredentials(awsID, awsKey, ""),
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

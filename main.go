package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv("ACCESS_TOKEN"),
		OAuthTokenSecret:     os.Getenv("ACCESS_TOKEN_SECRET"),
	}

	c, err := gotwi.NewClient(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	const DaysTotal int = 90
	var remainingDays uint = 90
	challenge := "#90DaysOfDevOps"

	fmt.Printf("Welcome to the %v challenge.\nThis challenge consists of %v days\n", challenge, DaysTotal)

	var TwitterName string
	var DaysCompleted uint

	// asking for user input
	fmt.Println("Enter Your Twitter Handle: ")
	fmt.Scanln(&TwitterName)

	fmt.Println("How many days have you completed?: ")
	fmt.Scanln(&DaysCompleted)

	// calculate remaining days
	remainingDays = remainingDays - DaysCompleted

	//fmt.Printf("Thank you %v for taking part and completing %v days.\n", TwitterName, DaysCompleted)
	//fmt.Printf("You have %v days remaining for the %v challenge\n", remainingDays, challenge)
	// fmt.Println("Good luck")

	text := fmt.Sprintf("Hey I am %v I have been doing the %v for %v days and I have %v Days left", TwitterName, challenge, DaysCompleted, remainingDays)

	p := &types.CreateInput{
		Text: gotwi.String(text),
		//Poll: &types.CreateInputPoll{
		//	DurationMinutes: gotwi.Int(5),
		//	Options: []string{
		//		"Cyan",
		//		"Magenta",
		//		"Yellow",
		//		"Key plate",
		//	},
		//},
	}

	res, err := managetweet.Create(context.Background(), c, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("[%s] %s\n", gotwi.StringValue(res.Data.ID), gotwi.StringValue(res.Data.Text))
}

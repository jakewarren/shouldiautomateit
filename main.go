package main

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey"
)

const (
	oneSecond         string = "1 second"
	fiveSeconds       string = "5 seconds"
	twentyfiveSeconds string = "25 seconds"
	thirtySeconds     string = "30 seconds"

	oneMinute         string = "1 minute"
	twoMinutes        string = "2 minutes"
	fourMinutes       string = "4 minutes"
	fiveMinutes       string = "5 minutes"
	twentyoneMinutes  string = "21 minutes"
	twentyfiveMinutes string = "25 minutes"
	thirtyMinutes     string = "30 minutes"

	oneHour        string = "1 hour"
	twoHours       string = "2 hours"
	fourHours      string = "4 hours"
	fiveHours      string = "5 hours"
	sixHours       string = "6 hours"
	twelveHours    string = "12 hours"
	twentyoneHours string = "21 hours"

	oneDay    string = "1 day"
	twoDays   string = "2 days"
	threeDays string = "3 days"
	sixDays   string = "6 days"
	fiveDays  string = "5 days"
	tenDays   string = "10 days"

	twoWeeks   string = "2 weeks"
	fourWeeks  string = "4 weeks"
	fiveWeeks  string = "5 weeks"
	eightWeeks string = "8 weeks"

	twoMonths  string = "2 months"
	sixMonths  string = "6 months"
	nineMonths string = "9 months"
	tenMonths  string = "10 months"
)

func main() {

	// the questions to ask
	var qs = []*survey.Question{
		{
			Name: "howOften",
			Prompt: &survey.Select{
				Message: "How often do you perform the task?",
				Options: []string{"50/day", "5/day", "Daily", "Weekly", "Monthly", "Yearly"},
			},
		},
		{
			Name: "timeSaved",
			Prompt: &survey.Select{
				Message: "How much time will you save by automating?",
				Options: []string{"1 second", "5 seconds", "30 seconds", "1 minute", "5 minutes", "30 minutes", "1 hour", "6 hours", "1 day"},
			},
		},
	}

	// the answers will be written to this struct
	answers := struct {
		HowOften  string
		TimeSaved string
	}{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	timeAllowed := allowedTime(strings.ToLower(answers.HowOften), answers.TimeSaved)
	fmt.Printf("You can spend up to %s automating the task assuming you will be performing the task over 5 years.\n", timeAllowed)

}

func allowedTime(howOften string, timeSaved string) string {

	switch howOften {
	case "50/day":
		switch timeSaved {
		case oneSecond:
			return oneDay
		case fiveSeconds:
			return fiveDays
		case thirtySeconds:
			return fourWeeks
		case oneMinute:
			return eightWeeks
		case fiveMinutes:
			return nineMonths
		}
	case "5/day":
		switch timeSaved {
		case oneSecond:
			return twoHours
		case fiveSeconds:
			return twelveHours
		case thirtySeconds:
			return threeDays
		case oneMinute:
			return sixDays
		case fiveMinutes:
			return fourWeeks
		case thirtyMinutes:
			return sixMonths
		case oneHour:
			return tenMonths
		}
	case "daily":
		switch timeSaved {
		case oneSecond:
			return thirtyMinutes
		case fiveSeconds:
			return twoHours
		case thirtySeconds:
			return twelveHours
		case oneMinute:
			return oneDay
		case fiveMinutes:
			return sixDays
		case thirtyMinutes:
			return fiveWeeks
		case oneHour:
			return twoMonths
		}
	case "weekly":
		switch timeSaved {
		case oneSecond:
			return fourMinutes
		case fiveSeconds:
			return twentyoneMinutes
		case thirtySeconds:
			return twoHours
		case oneMinute:
			return fourHours
		case fiveMinutes:
			return twentyoneHours
		case thirtyMinutes:
			return fiveDays
		case oneHour:
			return tenDays
		case sixHours:
			return twoMonths
		}
	case "monthly":
		switch timeSaved {
		case oneSecond:
			return oneMinute
		case fiveSeconds:
			return fiveMinutes
		case thirtySeconds:
			return thirtyMinutes
		case oneMinute:
			return oneHour
		case fiveMinutes:
			return fiveHours
		case thirtyMinutes:
			return oneDay
		case oneHour:
			return twoDays
		case sixHours:
			return twoWeeks
		case oneDay:
			return eightWeeks
		}
	case "yearly":
		switch timeSaved {
		case oneSecond:
			return fiveSeconds
		case fiveSeconds:
			return twentyfiveSeconds
		case thirtySeconds:
			return twoMinutes
		case oneMinute:
			return fiveMinutes
		case fiveMinutes:
			return twentyfiveMinutes
		case thirtyMinutes:
			return twoHours
		case oneHour:
			return fiveHours
		case sixHours:
			return oneDay
		case oneDay:
			return fiveDays
		}
	}

	return "not found"
}

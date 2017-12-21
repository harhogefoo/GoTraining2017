package main

import (
	"os"
	"log"
	"fmt"
	"github.com/harhogefoo/go_training2017/ch04/github"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	var monthIssues []*github.Issue
	var yearIssues []*github.Issue
	var otherIssues []*github.Issue
	now := time.Now()

	for _, item := range result.Items {
		if LessThanMonth(item.CreatedAt, now) {
			monthIssues = append(monthIssues, item)
			continue
		}

		if LessThanYear(item.CreatedAt, now) {
			yearIssues = append(yearIssues, item)
			continue
		}

		otherIssues = append(otherIssues, item)

		showIssues(monthIssues, "一ヶ月未満")
		showIssues(yearIssues, "一年未満")
		showIssues(otherIssues, "一年以上")
	}
}

func showIssues(issues []*github.Issue, header string) {
	if len(issues) < 0 {
		return
	}

	fmt.Printf("%s\n", header)
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func LessThanMonth(t, now time.Time) bool {
	month := now.AddDate(0, -1, 0)
	return t.After(month)
}

func LessThanYear(t, now time.Time) bool {
	year := now.AddDate(-1, 0, 0)
	return t.After(year)
}
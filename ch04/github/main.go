package github

import "time"

const IssueURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"created_at"`
	Body string // マークダウン(Markdown)形式
}

type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}

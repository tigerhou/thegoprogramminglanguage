package github

import "time"

const SearchURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:total_count`
	Items      []*Issue
}
type Issue struct {
	Number    int
	HTMLURL   string `html_url`
	State     string
	User      *User
	Title     string
	Body      string
	CreatedAt time.Time `created_at`
}
type User struct {
	Login   string
	HTMLURL string `html_url`
}

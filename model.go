package main

import "time"

type Project struct {
	ID string `json:"id"`
}

type Repository struct {
	Name   string `json:"name"`
	Commit string `json:"commit"`
}

type Document struct {
	ShortDescription string `json:"shortDescription"`
	URL              string `json:"url"`
}

type Scorecard struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
}

type Check struct {
	Name          string        `json:"name"`
	Documentation Document      `json:"documentation"`
	Score         int           `json:"score"`
	Reason        string        `json:"reason"`
	Details       []string      `json:"details"`
	OverallScore  float64       `json:"overallScore"`
	Metadata      []interface{} `json:"metadata"`
}

type Score struct {
	Date      time.Time  `json:"date"`
	Repo      Repository `json:"repository"`
	ScoreCard Scorecard  `json:"scorecard"`
	Checks    []Check    `json:"checks"`
}

type Response struct {
	ProjectKey      Project `json:"projectKey"`
	OpenIssuesCount int     `json:"openIssuesCount"`
	StarsCount      int     `json:"starsCount"`
	ForksCount      int     `json:"forksCount"`
	License         string  `json:"license"`
	Description     string  `json:"description"`
	Homepage        string  `json:"homepage"`
	ScoreCard       Score   `json:"scorecard"`
}

type Row struct {
	Name     string
	Json     string
	Packages string
}

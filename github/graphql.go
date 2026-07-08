package github

import "github.com/shurcooL/githubv4"

// Q = query

type RateLimit struct {
	Limit     githubv4.Int
	Remaining githubv4.Int
	ResetAt   githubv4.DateTime
}

type PageInfo struct {
	EndCursor   githubv4.String
	HasNextPage githubv4.Boolean
}

type QPing struct {
	Viewer struct {
		Login githubv4.String
	}
	RateLimit RateLimit
}

type QueryTeams struct {
	Organization struct {
		Teams struct {
			Nodes []struct {
				Slug githubv4.String
			}
			PageInfo PageInfo
		} `graphql:"teams(first: 100, after: $teamsCursor)"`
	} `graphql:"organization(login: $orgName)"`
}

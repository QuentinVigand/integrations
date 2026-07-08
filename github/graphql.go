package github

import "github.com/shurcooL/githubv4"

type RateLimit struct {
	Limit     githubv4.Int
	Remaining githubv4.Int
	ResetAt   githubv4.DateTime
}

type PageInfo struct {
	EndCursor   githubv4.String
	HasNextPage githubv4.Boolean
}

type QueryPing struct {
	Viewer struct {
		Login githubv4.String
	}
	RateLimit RateLimit
}

type QueryTeams struct {
	Organization struct {
		Teams struct {
			Nodes    []NodeTeam
			PageInfo PageInfo
		} `graphql:"teams(first: 100, after: $teamsCursor)"`
	} `graphql:"organization(login: $orgName)"`
}

type QueryMembers struct {
	Organization struct {
		Team struct {
			Name    githubv4.String
			Members struct {
				Nodes    []NodeMember
				PageInfo PageInfo
			} `graphql:"members(first: 100, after: $membersCursor)"` // Paginate members here
		} `graphql:"team(slug: $teamSlug)"`
	} `graphql:"organization(login: $orgName)"`
}
type NodeTeam struct {
	Slug githubv4.String
}
type NodeMember struct {
	Login githubv4.String
	Name  githubv4.String
}

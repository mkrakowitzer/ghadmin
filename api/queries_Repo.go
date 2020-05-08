package api

import (
	"github.com/shurcooL/githubv4"
	"github.com/spf13/cobra"
)

type ReposPayload struct {
	Organization struct {
		Repositories struct {
			Nodes []struct {
				Name string
				Id   string
			}
			PageInfo struct {
				HasNextPage bool
				EndCursor   githubv4.String
			}
		}
	}
}

func RepoList(client *Client, cmd *cobra.Command) (*ReposPayload, error) {
	query := `
	query($org: String!, $cursor: String) {
		organization(login: $org) {
		  repositories(first: 100,after: $cursor) {
			nodes {
			  name
			  id
			}
			pageInfo {
			  hasNextPage
			  endCursor
			}
		  }
		}
	}`
	org, _ := cmd.Flags().GetString("org")
	result := ReposPayload{}
	allResults := ReposPayload{}
	variables := map[string]interface{}{
		"org":    githubv4.String(org),
		"cursor": (*githubv4.String)(nil),
	}
	for {
		err := client.GraphQL(query, variables, &result)
		if err != nil {
			return nil, err
		}
		allResults.Organization.Repositories.Nodes = append(allResults.Organization.Repositories.Nodes, result.Organization.Repositories.Nodes...)
		if !result.Organization.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(result.Organization.Repositories.PageInfo.EndCursor)
	}
	return &allResults, nil
}

type RepoPayload struct {
	Organization struct {
		Repository struct {
			ID githubv4.ID
		}
	}
}

func GetRepoID(client *Client, cmd *cobra.Command, args string) (*RepoPayload, error) {
	query := `
	query($org: String!, $name: String!) {
	  organization(login: $org) {
	    repository(name: $name) {
	      id
	    }
	  }
	}`
	org, _ := cmd.Flags().GetString("org")
	variables := map[string]interface{}{"org": org, "name": args}
	result := RepoPayload{}

	err := client.GraphQL(query, variables, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//hasProjectsEnabled: false, hasWikiEnabled: false
func DisableIssues(client *Client, cmd *cobra.Command, variables map[string]interface{}) error {
	mutation := `
	mutation ($id: ID!, $issues: Boolean!) {
		updateRepository(input: {hasIssuesEnabled: $issues, repositoryId: $id}) {
		  clientMutationId
		}
	  }`
	err := client.GraphQL(mutation, variables, nil)
	if err != nil {
		return err
	}
	return nil
}

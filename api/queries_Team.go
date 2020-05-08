package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/shurcooL/githubv4"
	"github.com/spf13/cobra"
)

func TeamDelete(client *Client, cmd *cobra.Command, args []string) (string, error) {

	name, _ := cmd.Flags().GetString("name")
	org, _ := cmd.Flags().GetString("org")
	path := fmt.Sprintf("orgs/%s/teams/%s", org, name)

	r := bytes.NewReader([]byte(`{}`))
	err := client.REST("DELETE", path, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	return name, err
}

type Team struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Permission  string `json:"permission"`
	Privacy     string `json:"privacy"`
}

func TeamCreate(client *Client, cmd *cobra.Command, args []string) (*Team, error) {

	name, _ := cmd.Flags().GetString("name")
	org, _ := cmd.Flags().GetString("org")
	desc, _ := cmd.Flags().GetString("desc")
	path := fmt.Sprintf("orgs/%s/teams", org)
	result := Team{}

	team := Team{
		Name:        name,
		Description: desc,
		Permission:  "pull",
		Privacy:     "closed",
	}

	j, _ := json.Marshal(team)

	err := client.REST("POST", path, bytes.NewBuffer(j), &result)

	if err != nil {
		log.Fatal(err)
	}

	return &Team{
		Name:        result.Name,
		Description: result.Description,
		Permission:  result.Permission,
		Privacy:     result.Privacy,
	}, nil
}

type Teams struct {
	Organization struct {
		Teams struct {
			Edges []struct {
				Node struct {
					Name githubv4.String
				}
			}
		}
	}
}

func TeamList(client *Client, cmd *cobra.Command) (*Teams, error) {

	// If you have more than a 100 teams you have bigger problems.
	query := `
	query($org: String!) {
		organization(login: $org ) {
		  teams(first: 100) {
			edges {
			  node {
				name
			  }
			}
		  }
		}
	}`
	org, _ := cmd.Flags().GetString("org")
	variables := map[string]interface{}{"org": org}
	result := Teams{}

	err := client.GraphQL(query, variables, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil

}

type TeamMembers struct {
	Organization struct {
		Team struct {
			Members struct {
				Nodes []struct {
					Login      string
					Name       string
					DatabaseID int
					Location   string
				}
				TotalCount int
			}
		}
	}
}

func TeamListMembers(client *Client, cmd *cobra.Command, name string) (*TeamMembers, error) {

	// If you have more than a 100 teams you have bigger problems.
	query := `
	query($org: String!, $name: String!) {
		organization(login: $org) {
		  team(slug: $name) {
			members {
			  nodes {
				login
				name
				databaseId
			  }
			  totalCount
			}
		  }
		}
	  }`

	org, _ := cmd.Flags().GetString("org")
	variables := map[string]interface{}{"org": org, "name": name}
	result := TeamMembers{}

	err := client.GraphQL(query, variables, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil

}

func TeamAddToRepo(client *Client, cmd *cobra.Command, args []string) error {

	name, _ := cmd.Flags().GetString("name")
	org, _ := cmd.Flags().GetString("org")
	repo, _ := cmd.Flags().GetString("repo")
	permission, _ := cmd.Flags().GetString("permission")
	path := fmt.Sprintf("orgs/%s/teams/%s/repos/%s/%s", org, name, org, repo)
	result := Team{}

	team := Team{
		Permission: permission,
	}

	j, _ := json.Marshal(team)

	err := client.REST("PUT", path, bytes.NewBuffer(j), &result)

	if err != nil {
		log.Fatal(err)
	}
	return nil
}

type Role struct {
	URL   string
	Role  string
	State string
}

func TeamMemberships(client *Client, cmd *cobra.Command, team string) (*Role, error) {
	username, _ := cmd.Flags().GetString("username")
	org, _ := cmd.Flags().GetString("org")
	r, _ := cmd.Flags().GetString("role")
	f, _ := cmd.Flags().GetBool("delete")

	var method string
	if f {
		method = "DELETE"
	} else {
		method = "PUT"
	}

	path := fmt.Sprintf("orgs/%s/teams/%s/memberships/%s", org, team, username)
	result := Role{}

	role := Role{
		Role: r,
	}

	j, _ := json.Marshal(role)

	err := client.REST(method, path, bytes.NewBuffer(j), &result)
	if err != nil {
		log.Fatal(err)
	}

	return &result, nil
}

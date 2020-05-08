package cmd

import (
	"fmt"

	"github.com/krakowitzerm/ghadmin/api"
	"github.com/spf13/cobra"
)

var repoListCmd = &cobra.Command{
	Use:   "list",
	Short: "List repositories in a GitHub Org",
	RunE:  repoList,
}

func init() {
	repoCmd.AddCommand(repoListCmd)
}

func repoList(cmd *cobra.Command, args []string) error {
	ctx := contextForCommand(cmd)
	apiClient, err := apiClientForContext(ctx)
	if err != nil {
		return err
	}

	Payload, err := api.RepoList(apiClient, cmd)
	if err != nil {
		return err
	}

	out := colorableOut(cmd)
	for _, repo := range Payload.Organization.Repositories.Nodes {
		fmt.Fprintln(out, repo.Name)
	}
	return nil
}

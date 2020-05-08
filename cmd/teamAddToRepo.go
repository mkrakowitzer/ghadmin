package cmd

import (
	"fmt"

	"github.com/krakowitzerm/ghadmin/api"
	"github.com/krakowitzerm/ghadmin/utils"
	"github.com/spf13/cobra"
)

var teamAddToRepoCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a team to a repo",
	Args:  cobra.ExactArgs(0),
	RunE:  teamAddToRepo,
}

func init() {
	teamCmd.AddCommand(teamAddToRepoCmd)
	teamAddToRepoCmd.Flags().StringP("repo", "r", "", "The repository to add this team to")
	teamAddToRepoCmd.Flags().StringP("permission", "p", "pull", "The permission to assign this team")
}

func teamAddToRepo(cmd *cobra.Command, args []string) error {
	ctx := contextForCommand(cmd)
	apiClient, err := apiClientForContext(ctx)
	if err != nil {
		return err
	}

	err = api.TeamAddToRepo(apiClient, cmd, args)
	if err != nil {
		return err
	}
	repo, _ := cmd.Flags().GetString("repo")
	name, _ := cmd.Flags().GetString("name")
	permission, _ := cmd.Flags().GetString("permission")

	greenCheck := utils.Green("✓")
	fmt.Printf("%s Added team %s to %s with permissions %s\n", greenCheck, name, repo, permission)

	return nil
}

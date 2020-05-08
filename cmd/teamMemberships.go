package cmd

import (
	"fmt"

	"github.com/krakowitzerm/ghadmin/api"
	"github.com/krakowitzerm/ghadmin/utils"
	"github.com/spf13/cobra"
)

var teamMembershipsCmd = &cobra.Command{
	Use:   "memberships",
	Short: "Manage team memberships",
	Args:  cobra.ExactArgs(1),
	RunE:  teamMemberships,
}

func init() {
	teamCmd.AddCommand(teamMembershipsCmd)
	teamMembershipsCmd.Flags().StringP("role", "r", "member", "The role of the member")
	teamMembershipsCmd.Flags().StringP("username", "u", "", "The username")
	teamMembershipsCmd.Flags().BoolP("delete", "d", false, "All repositories")
}

func teamMemberships(cmd *cobra.Command, args []string) error {
	ctx := contextForCommand(cmd)
	apiClient, err := apiClientForContext(ctx)
	if err != nil {
		return err
	}

	team, err := api.TeamMemberships(apiClient, cmd, args[0])
	if err != nil {
		return err
	}

	username, _ := cmd.Flags().GetString("username")
	greenCheck := utils.Green("✓")
	fmt.Printf("%s Added %s to %s with role %s\n", greenCheck, username, args[0], team.Role)

	return nil
}

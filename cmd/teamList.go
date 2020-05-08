package cmd

import (
	"fmt"

	"github.com/krakowitzerm/ghadmin/api"
	"github.com/spf13/cobra"
)

var teamListCmd = &cobra.Command{
	Use:   "list",
	Short: "List teams",
	Args:  cobra.ExactArgs(0),
	RunE:  teamList,
}

func init() {
	teamCmd.AddCommand(teamListCmd)
}

func teamList(cmd *cobra.Command, args []string) error {
	ctx := contextForCommand(cmd)
	apiClient, err := apiClientForContext(ctx)
	if err != nil {
		return err
	}
	out := colorableOut(cmd)
	Payload, err := api.TeamList(apiClient, cmd)
	if err != nil {
		return err
	}

	for _, team := range Payload.Organization.Teams.Edges {
		fmt.Fprintln(out, team.Node.Name)
	}
	return nil
}

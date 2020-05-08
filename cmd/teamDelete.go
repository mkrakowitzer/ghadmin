package cmd

import (
	"fmt"

	"github.com/krakowitzerm/ghadmin/api"
	"github.com/krakowitzerm/ghadmin/utils"
	"github.com/spf13/cobra"
)

var teamDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a Github team",
	Args:  cobra.ExactArgs(0),
	RunE:  teamDelete,
}

func init() {
	teamCmd.AddCommand(teamDeleteCmd)
	teamListCmd.PersistentFlags().StringP("name", "n", "", "The team name")
}

func teamDelete(cmd *cobra.Command, args []string) error {
	ctx := contextForCommand(cmd)
	apiClient, err := apiClientForContext(ctx)
	if err != nil {
		return err
	}

	name, err := api.TeamDelete(apiClient, cmd, args)
	if err != nil {
		return err
	}

	greenCheck := utils.Green("✓")
	fmt.Printf("%s Deleted team %s\n", greenCheck, name)

	return nil
}

package cmd

import (
	"fmt"

	"github.com/krakowitzerm/ghadmin/api"
	"github.com/krakowitzerm/ghadmin/utils"
	"github.com/spf13/cobra"
)

var teamCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Github team",
	Args:  cobra.ExactArgs(0),
	RunE:  teamCreate,
}

func init() {
	teamCmd.AddCommand(teamCreateCmd)
	teamCreateCmd.Flags().StringP("desc", "d", "", "Description")
	teamCreateCmd.PersistentFlags().StringP("name", "n", "", "The team name")
}

func teamCreate(cmd *cobra.Command, args []string) error {
	ctx := contextForCommand(cmd)
	apiClient, err := apiClientForContext(ctx)
	if err != nil {
		return err
	}

	team, err := api.TeamCreate(apiClient, cmd, args)
	if err != nil {
		return err
	}

	greenCheck := utils.Green("✓")
	fmt.Printf("%s Created team %s\n", greenCheck, team.Name)

	return nil
}

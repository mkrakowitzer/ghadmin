package cmd

import (
	"os"
	"strconv"

	"github.com/krakowitzerm/ghadmin/api"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var teamListMembersCmd = &cobra.Command{
	Use:   "members",
	Short: "List team members",
	Args:  cobra.ExactArgs(1),
	RunE:  teamListMembers,
}

func init() {
	teamListCmd.AddCommand(teamListMembersCmd)
}

func teamListMembers(cmd *cobra.Command, args []string) error {
	ctx := contextForCommand(cmd)
	apiClient, err := apiClientForContext(ctx)
	if err != nil {
		return err
	}

	Payload, err := api.TeamListMembers(apiClient, cmd, args[0])
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Login", "Name"})
	table.SetFooter([]string{"Total", strconv.Itoa(Payload.Organization.Team.Members.TotalCount)})
	table.SetFooterColor(tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.FgHiRedColor})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetFooterAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("  ")
	table.SetColumnSeparator("  ")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	for _, v := range Payload.Organization.Team.Members.Nodes {
		table.Append([]string{v.Login, v.Name})
	}
	table.Render()

	return nil
}

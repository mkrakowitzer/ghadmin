package cmd

import (
	"fmt"
	"log"

	"github.com/krakowitzerm/ghadmin/api"
	"github.com/krakowitzerm/ghadmin/utils"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a repositories properties",
	Long:  `Edit repository properties`,
	RunE:  repoEdit,
}

func repoEdit(cmd *cobra.Command, args []string) error {

	allStatus, _ := cmd.Flags().GetBool("all")
	if len(args) > 1 && allStatus {
		log.Fatal("Can not specify --all with an argument for repo")
	}

	issueStatus, _ := cmd.Flags().GetBool("disable-issues")

	ctx := contextForCommand(cmd)
	apiClient, err := apiClientForContext(ctx)
	if err != nil {
		return err
	}

	if issueStatus {
		err := disableIssue(apiClient, cmd, args, allStatus)
		if err != nil {
			return err
		}
	}
	greenCheck := utils.Green("✓")
	fmt.Printf("%s Edit Done\n", greenCheck)
	return nil
}

func init() {

	repoCmd.AddCommand(editCmd)
	editCmd.Flags().BoolP("disable-issues", "", false, "Disable Issues")
}

func disableIssue(apiClient *api.Client, cmd *cobra.Command, args []string, allStatus bool) error {

	if allStatus {
		Payload, err := api.RepoList(apiClient, cmd)
		if err != nil {
			return err
		}
		for _, repo := range Payload.Organization.Repositories.Nodes {
			variables := map[string]interface{}{"id": repo.Id, "issues": false}
			err = api.DisableIssues(apiClient, cmd, variables)
			if err != nil {
				return err
			}
		}
	} else {
		repoId, err := api.GetRepoID(apiClient, cmd, args[0])
		if err != nil {
			return err
		}
		variables := map[string]interface{}{"id": repoId.Organization.Repository.ID, "issues": false}
		err = api.DisableIssues(apiClient, cmd, variables)
		if err != nil {
			return err
		}
	}
	return nil
}

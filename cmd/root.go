package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/krakowitzerm/ghadmin/api"
	"github.com/krakowitzerm/ghadmin/context"
	"github.com/krakowitzerm/ghadmin/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ghadmin",
	Short: "CLI tool for managing a GitHub Organisation",
	Long: `CLI tool for managing a GitHub Organisation

You should set an environment variable GITHUB_TOKEN for authenticated requests
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringP("org", "o", "", "The GitHub organisation name")
}

func contextForCommand(cmd *cobra.Command) context.Context {
	ctx := initContext()
	return ctx
}

// overridden in tests
var initContext = func() context.Context {
	ctx := context.New()
	return ctx
}

var apiClientForContext = func(ctx context.Context) (*api.Client, error) {
	token, err := ctx.AuthToken()
	if err != nil {
		return nil, err
	}

	var opts []api.ClientOption

	getAuthValue := func() string {
		return fmt.Sprintf("token %s", token)
	}

	Version := "1"
	opts = append(opts,
		//api.CheckScopes("read:org", checkScopesFunc),
		api.AddHeaderFunc("Authorization", getAuthValue),
		api.AddHeader("User-Agent", fmt.Sprintf("ghadmin %s", Version)),
		// antiope-preview: Checks
		api.AddHeader("Accept", "application/vnd.github.antiope-preview+json"),
	)

	return api.NewClient(opts...), nil

}

func colorableOut(cmd *cobra.Command) io.Writer {
	out := cmd.OutOrStdout()
	if outFile, isFile := out.(*os.File); isFile {
		return utils.NewColorable(outFile)
	}
	return out
}

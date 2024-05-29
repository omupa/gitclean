package command

import (
	"fmt"
	"os"

	"github.com/omupa/gitclean/interative"
	"github.com/omupa/gitclean/operator"
	"github.com/spf13/cobra"
)

func buildCommands() *cobra.Command {
	var forceFlag bool
	var deleteAllFlag bool

	rootCmd := &cobra.Command{
		Use:   "gitclean",
		Short: "Gitclean is a tools to exclude something from git repository",
		Long:  "With gitclean we can exclude: branches, coming soon",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			operator.HasGit()

			if deleteAllFlag {
				operator.DeleteAll(forceFlag)
				return
			}

			interative.StartInterative(forceFlag)
		},
	}
	rootCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Force deletion like: git branch -D branch_name")
	rootCmd.Flags().BoolVarP(&deleteAllFlag, "all", "a", false, "Force deletion for all. It will not start interative mode")

	return rootCmd
}

func Execute() {
	if err := buildCommands().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

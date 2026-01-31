package rssaurus

import (
	"fmt"
	"os"

	"github.com/justinburdett/rssaurus-cli/internal/api"
	"github.com/justinburdett/rssaurus-cli/internal/config"
	"github.com/spf13/cobra"
)

var (
	flagHost   string
	flagJSON   bool
	cfgManager *config.Manager
	apiClient  *api.Client
)

var rootCmd = &cobra.Command{
	Use:   "rssaurus",
	Short: "RSSaurus command-line client",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		m, err := config.NewManager()
		if err != nil {
			return err
		}
		cfgManager = m

		// Allow overrides
		if flagHost != "" {
			cfgManager.SetHost(flagHost)
		}
		if os.Getenv("RSSAURUS_HOST") != "" {
			cfgManager.SetHost(os.Getenv("RSSAURUS_HOST"))
		}
		if os.Getenv("RSSAURUS_TOKEN") != "" {
			cfgManager.SetToken(os.Getenv("RSSAURUS_TOKEN"))
		}

		apiClient = api.NewClient(cfgManager.Host(), cfgManager.Token())

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&flagHost, "host", "", "RSSaurus host (default https://rssaurus.com)")
	rootCmd.PersistentFlags().BoolVar(&flagJSON, "json", false, "output JSON")

	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(feedsCmd)
	rootCmd.AddCommand(itemsCmd)
	rootCmd.AddCommand(openCmd)
	rootCmd.AddCommand(readCmd)
	rootCmd.AddCommand(unreadCmd)
	rootCmd.AddCommand(markReadCmd)
	rootCmd.AddCommand(saveCmd)
	rootCmd.AddCommand(unsaveCmd)
}

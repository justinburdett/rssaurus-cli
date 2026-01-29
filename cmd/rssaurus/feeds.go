package rssaurus

import (
	"fmt"
	"os"

	"github.com/justinburdett/rssaurus-cli/internal/api"
	"github.com/justinburdett/rssaurus-cli/internal/output"
	"github.com/spf13/cobra"
)

var feedsCmd = &cobra.Command{
	Use:   "feeds",
	Short: "List feeds",
	RunE: func(cmd *cobra.Command, args []string) error {
		var resp api.FeedsResponse
		if err := apiClient.GetJSON(cmd.Context(), "/api/v1/feeds", &resp); err != nil {
			return err
		}

		if flagJSON {
			return output.PrintJSON(os.Stdout, resp)
		}

		tw := output.NewTabWriter(os.Stdout)
		defer tw.Flush()

		fmt.Fprintln(tw, "ID\tUNREAD\tTITLE\tFEED_URL")
		for _, f := range resp.Feeds {
			fmt.Fprintf(tw, "%d\t%d\t%s\t%s\n", f.ID, f.UnreadCount, output.Trunc(f.Title, 60), output.Trunc(f.FeedURL, 80))
		}
		return nil
	},
}

package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/justinburdett/rssaurus-cli/internal/api"
	"github.com/justinburdett/rssaurus-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	itemsFeedID int64
	itemsStatus string
	itemsLimit  int
	itemsCursor string
	itemsURLs   bool
)

var itemsCmd = &cobra.Command{
	Use:   "items",
	Short: "List feed items (unread by default)",
	RunE: func(cmd *cobra.Command, args []string) error {
		q := url.Values{}
		if itemsFeedID > 0 {
			q.Set("feed_id", strconv.FormatInt(itemsFeedID, 10))
		}
		if itemsStatus != "" {
			q.Set("status", itemsStatus)
		} else {
			q.Set("status", "unread")
		}
		if itemsLimit > 0 {
			q.Set("limit", strconv.Itoa(itemsLimit))
		}
		if itemsCursor != "" {
			q.Set("cursor", itemsCursor)
		}

		path := "/api/v1/items"
		if enc := q.Encode(); enc != "" {
			path += "?" + enc
		}

		var resp api.ItemsResponse
		if err := apiClient.GetJSON(cmd.Context(), path, &resp); err != nil {
			return err
		}

		if flagJSON {
			return output.PrintJSON(os.Stdout, resp)
		}

		if itemsURLs {
			for _, it := range resp.Items {
				if it.URL != "" {
					fmt.Fprintln(os.Stdout, it.URL)
				}
			}
			// No cursor printing in --urls mode (keeps piping clean).
			return nil
		}

		tw := output.NewTabWriter(os.Stdout)
		defer tw.Flush()

		fmt.Fprintln(tw, "STATUS\tPUBLISHED\tFEED\tTITLE\tURL")
		for _, it := range resp.Items {
			status := "unread"
			if it.ReadAt != nil {
				status = "read"
			}
			published := ""
			if it.PublishedAt != nil {
				published = it.PublishedAt.Local().Format("2006-01-02")
			}
			fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\n",
				status,
				published,
				output.Trunc(it.FeedTitle, 24),
				output.Trunc(it.Title, 70),
				output.Trunc(it.URL, 70),
			)
		}

		if resp.NextCursor != "" {
			fmt.Fprintf(os.Stderr, "Next cursor: %s\n", resp.NextCursor)
		}
		return nil
	},
}

func init() {
	itemsCmd.Flags().Int64Var(&itemsFeedID, "feed-id", 0, "filter by feed id")
	itemsCmd.Flags().StringVar(&itemsStatus, "status", "unread", "status filter: unread|read|all")
	itemsCmd.Flags().IntVar(&itemsLimit, "limit", 50, "items per page (max 200)")
	itemsCmd.Flags().StringVar(&itemsCursor, "cursor", "", "pagination cursor")
	itemsCmd.Flags().BoolVar(&itemsURLs, "urls", false, "print only item URLs (one per line)")
}

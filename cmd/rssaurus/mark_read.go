package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/justinburdett/rssaurus-cli/internal/api"
	"github.com/justinburdett/rssaurus-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	markReadAll    bool
	markReadFeedID int64
	markReadIDs    string
	markReadURL    string
	markReadURLs   string
)

var markReadCmd = &cobra.Command{
	Use:   "mark-read",
	Short: "Mark many items as read",
	RunE: func(cmd *cobra.Command, args []string) error {
		payload := map[string]any{}
		if markReadFeedID > 0 {
			payload["feed_id"] = markReadFeedID
		}
		if markReadAll {
			payload["all"] = true
		} else if strings.TrimSpace(markReadIDs) != "" {
			parts := strings.Split(markReadIDs, ",")
			ids := make([]int64, 0, len(parts))
			for _, p := range parts {
				p = strings.TrimSpace(p)
				if p == "" {
					continue
				}
				id, err := strconv.ParseInt(p, 10, 64)
				if err != nil {
					return fmt.Errorf("invalid id %q", p)
				}
				ids = append(ids, id)
			}
			payload["ids"] = ids
		} else if strings.TrimSpace(markReadURL) != "" {
			payload["url"] = strings.TrimSpace(markReadURL)
		} else if strings.TrimSpace(markReadURLs) != "" {
			parts := strings.Split(markReadURLs, ",")
			urls := make([]string, 0, len(parts))
			for _, p := range parts {
				p = strings.TrimSpace(p)
				if p == "" {
					continue
				}
				urls = append(urls, p)
			}
			payload["urls"] = urls
		} else {
			return fmt.Errorf("provide one of --all, --ids, --url, or --urls")
		}

		var resp api.MarkReadResponse
		if err := apiClient.PostJSON(cmd.Context(), "/api/v1/items/mark_read", payload, &resp); err != nil {
			return err
		}
		if flagJSON {
			return output.PrintJSON(os.Stdout, resp)
		}
		fmt.Printf("Updated %d item(s)\n", resp.Updated)
		return nil
	},
}

func init() {
	markReadCmd.Flags().BoolVar(&markReadAll, "all", false, "mark all items as read (optionally filtered by --feed-id)")
	markReadCmd.Flags().Int64Var(&markReadFeedID, "feed-id", 0, "filter by feed id")
	markReadCmd.Flags().StringVar(&markReadIDs, "ids", "", "comma-separated item ids")
	markReadCmd.Flags().StringVar(&markReadURL, "url", "", "mark the item with this URL as read")
	markReadCmd.Flags().StringVar(&markReadURLs, "urls", "", "comma-separated URLs to mark as read")
}

package rssaurus

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
	saveTitle string
)

var saveCmd = &cobra.Command{
	Use:   "save <url>",
	Short: "Save an article by URL",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		u := args[0]
		if _, err := url.ParseRequestURI(u); err != nil {
			return fmt.Errorf("invalid url")
		}
		payload := map[string]any{"url": u}
		if saveTitle != "" {
			payload["title"] = saveTitle
		}
		var resp api.SavedCreateResponse
		if err := apiClient.PostJSON(cmd.Context(), "/api/v1/saved_items", payload, &resp); err != nil {
			return err
		}
		if flagJSON {
			return output.PrintJSON(os.Stdout, resp)
		}
		fmt.Printf("Saved: %s (id=%d)\n", resp.URL, resp.ID)
		return nil
	},
}

var unsaveCmd = &cobra.Command{
	Use:   "unsave <saved-item-id>",
	Short: "Unsave (delete) a saved item by id",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid id")
		}
		if err := apiClient.Delete(cmd.Context(), fmt.Sprintf("/api/v1/saved_items/%d", id)); err != nil {
			return err
		}
		if flagJSON {
			fmt.Fprintln(os.Stdout, "{}")
			return nil
		}
		fmt.Printf("Unsaved: %d\n", id)
		return nil
	},
}

func init() {
	saveCmd.Flags().StringVar(&saveTitle, "title", "", "optional title")
}

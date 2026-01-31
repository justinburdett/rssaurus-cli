package rssaurus

import (
	"fmt"
	"os"
	"strconv"

	"github.com/justinburdett/rssaurus-cli/internal/api"
	"github.com/justinburdett/rssaurus-cli/internal/output"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read <item-id>",
	Short: "Mark an item as read",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid item id: %w", err)
		}
		var resp api.ReadStateResponse
		if err := apiClient.PostJSON(cmd.Context(), fmt.Sprintf("/api/v1/items/%d/read", id), map[string]any{}, &resp); err != nil {
			return err
		}
		if flagJSON {
			return output.PrintJSON(os.Stdout, resp)
		}
		fmt.Println("Marked read.")
		return nil
	},
}

var unreadCmd = &cobra.Command{
	Use:   "unread <item-id>",
	Short: "Mark an item as unread",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid item id: %w", err)
		}
		var resp api.ReadStateResponse
		if err := apiClient.PostJSON(cmd.Context(), fmt.Sprintf("/api/v1/items/%d/unread", id), map[string]any{}, &resp); err != nil {
			return err
		}
		if flagJSON {
			return output.PrintJSON(os.Stdout, resp)
		}
		fmt.Println("Marked unread.")
		return nil
	},
}

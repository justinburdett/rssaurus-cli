package rssaurus

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/justinburdett/rssaurus-cli/internal/platform"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open <url>",
	Short: "Open a URL in your default browser",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		raw := strings.TrimSpace(args[0])
		if raw == "" {
			return fmt.Errorf("url is required")
		}
		u, err := url.ParseRequestURI(raw)
		if err != nil || u.Scheme == "" || u.Host == "" {
			return fmt.Errorf("invalid url")
		}

		if err := platform.OpenURL(raw); err != nil {
			return err
		}

		// Keep output minimal; URL is useful confirmation without leaking IDs.
		if flagJSON {
			fmt.Fprintf(os.Stdout, "{\"opened\":%q}\n", raw)
			return nil
		}
		fmt.Fprintf(os.Stdout, "Opened: %s\n", raw)
		return nil
	},
}

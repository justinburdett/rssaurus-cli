package rssaurus

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication helpers",
}

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Open RSSaurus token creation page, then save a token locally",
	RunE: func(cmd *cobra.Command, args []string) error {
		createURL := "https://rssaurus.com/api_tokens/new"
		_ = openBrowser(createURL)

		fmt.Fprintf(os.Stderr, "Opened: %s\n", createURL)
		fmt.Fprint(os.Stderr, "Paste your API token: ")

		r := bufio.NewReader(os.Stdin)
		token, err := r.ReadString('\n')
		if err != nil {
			return err
		}
		token = strings.TrimSpace(token)
		if token == "" {
			return fmt.Errorf("token cannot be empty")
		}

		cfgManager.SetToken(token)
		if err := cfgManager.Save(); err != nil {
			return err
		}

		fmt.Fprintln(os.Stderr, "Saved token.")
		return nil
	},
}

func init() {
	authCmd.AddCommand(authLoginCmd)
}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	return cmd.Start()
}

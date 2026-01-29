package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Config struct {
	Host  string `json:"host"`
	Token string `json:"token"`
}

type Manager struct {
	cfg  Config
	path string
}

func NewManager() (*Manager, error) {
	path, err := defaultConfigPath()
	if err != nil {
		return nil, err
	}

	m := &Manager{
		cfg: Config{
			Host: "https://rssaurus.com",
		},
		path: path,
	}

	_ = m.Load() // best-effort
	return m, nil
}

func defaultConfigPath() (string, error) {
	// Prefer XDG-style paths so users can predict where config is stored.
	//
	// On macOS, os.UserConfigDir() points to ~/Library/Application Support,
	// which is fine, but many CLI users expect ~/.config.
	//
	// Order:
	// 1) $XDG_CONFIG_HOME
	// 2) ~/.config
	// 3) os.UserConfigDir() fallback
	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		return filepath.Join(xdg, "rssaurus", "config.json"), nil
	}

	home, err := os.UserHomeDir()
	if err == nil && home != "" {
		return filepath.Join(home, ".config", "rssaurus", "config.json"), nil
	}

	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "rssaurus", "config.json"), nil
}

func (m *Manager) Load() error {
	b, err := os.ReadFile(m.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	var c Config
	if err := json.Unmarshal(b, &c); err != nil {
		return err
	}
	if c.Host != "" {
		m.cfg.Host = c.Host
	}
	if c.Token != "" {
		m.cfg.Token = c.Token
	}
	return nil
}

func (m *Manager) Save() error {
	if err := os.MkdirAll(filepath.Dir(m.path), 0o755); err != nil {
		return err
	}

	b, err := json.MarshalIndent(m.cfg, "", "  ")
	if err != nil {
		return err
	}
	b = append(b, '\n')
	return os.WriteFile(m.path, b, 0o600)
}

func (m *Manager) Host() string  { return m.cfg.Host }
func (m *Manager) Token() string { return m.cfg.Token }

func (m *Manager) SetHost(host string)   { m.cfg.Host = host }
func (m *Manager) SetToken(token string) { m.cfg.Token = token }

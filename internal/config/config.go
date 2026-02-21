package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

// AgentMailConfig holds mcp_agent_mail integration settings.
type AgentMailConfig struct {
	Enabled bool `yaml:"enabled"`
	Port    int  `yaml:"port"`
}

// Config holds application configuration.
type Config struct {
	RefreshInterval time.Duration   `yaml:"refresh_interval"`
	SessionPrefix   string          `yaml:"session_prefix"`
	DefaultDir      string          `yaml:"default_dir"`
	DefaultArgs     string          `yaml:"default_args"`
	LogHistory      int             `yaml:"log_history"`
	AgentMail       AgentMailConfig `yaml:"agent_mail"`
}

// configFile is the YAML representation.
type configFile struct {
	RefreshInterval string          `yaml:"refresh_interval"`
	SessionPrefix   string          `yaml:"session_prefix"`
	DefaultDir      string          `yaml:"default_dir"`
	DefaultArgs     string          `yaml:"default_args"`
	LogHistory      int             `yaml:"log_history"`
	AgentMail       AgentMailConfig `yaml:"agent_mail"`
}

// DefaultConfig returns the default configuration.
func DefaultConfig() *Config {
	return &Config{
		RefreshInterval: 2 * time.Second,
		SessionPrefix:   "cd-",
		DefaultDir:      "",
		LogHistory:      1000,
		AgentMail: AgentMailConfig{
			Enabled: false,
			Port:    8765,
		},
	}
}

// ConfigDir returns the config directory path.
func ConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: could not determine home directory: %v\n", err)
	}
	return filepath.Join(home, ".claude-dashboard")
}

// ConfigPath returns the config file path.
func ConfigPath() string {
	return filepath.Join(ConfigDir(), "config.yaml")
}

// Load reads configuration from file, falling back to defaults.
func Load() *Config {
	cfg := DefaultConfig()

	data, err := os.ReadFile(ConfigPath())
	if err != nil {
		return cfg
	}

	var cf configFile
	if err := yaml.Unmarshal(data, &cf); err != nil {
		return cfg
	}

	if cf.RefreshInterval != "" {
		if d, err := time.ParseDuration(cf.RefreshInterval); err == nil {
			cfg.RefreshInterval = d
		}
	}
	if cf.SessionPrefix != "" {
		cfg.SessionPrefix = cf.SessionPrefix
	}
	if cf.DefaultDir != "" {
		cfg.DefaultDir = cf.DefaultDir
	}
	if cf.DefaultArgs != "" {
		cfg.DefaultArgs = cf.DefaultArgs
	}
	if cf.LogHistory > 0 {
		cfg.LogHistory = cf.LogHistory
	}
	if cf.AgentMail.Port != 0 {
		cfg.AgentMail.Port = cf.AgentMail.Port
	}
	cfg.AgentMail.Enabled = cf.AgentMail.Enabled

	return cfg
}

// Save writes the configuration to file.
func Save(cfg *Config) error {
	dir := ConfigDir()
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	cf := configFile{
		RefreshInterval: cfg.RefreshInterval.String(),
		SessionPrefix:   cfg.SessionPrefix,
		DefaultDir:      cfg.DefaultDir,
		DefaultArgs:     cfg.DefaultArgs,
		LogHistory:      cfg.LogHistory,
		AgentMail:       cfg.AgentMail,
	}

	data, err := yaml.Marshal(&cf)
	if err != nil {
		return err
	}

	return os.WriteFile(ConfigPath(), data, 0644)
}

package main

import (
	"fmt"
	"os"

	"github.com/seunggabi/claude-dashboard/internal/app"
)

var version = "dev"

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--version", "-v":
			fmt.Printf("claude-dashboard %s\n", version)
			os.Exit(0)
		case "--help", "-h":
			printHelp()
			os.Exit(0)
		case "attach":
			if len(os.Args) < 3 {
				fmt.Fprintln(os.Stderr, "Usage: claude-dashboard attach <session-name>")
				os.Exit(1)
			}
			if err := app.ExecAttach(os.Args[2]); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			os.Exit(0)
		}
	}

	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println(`claude-dashboard - k9s-style Claude Code Session Manager

Usage:
  claude-dashboard              Start the TUI dashboard
  claude-dashboard attach NAME  Attach to a session directly
  claude-dashboard --version    Show version
  claude-dashboard --help       Show this help

Keybindings:
  enter   Attach to session
  n       New session
  K       Kill session
  l       View logs
  d       Session detail
  /       Filter
  r       Refresh
  ?       Help
  q       Quit

Requirements:
  - tmux must be installed

Config:
  ~/.claude-dashboard/config.yaml`)
}

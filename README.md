# claude-dashboard

k9s-style TUI for managing Claude Code sessions via tmux.

![Go](https://img.shields.io/badge/Go-1.25-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Features

- **Session Dashboard** - View all Claude Code tmux sessions in a table
- **Attach/Detach** - Enter to attach, `Ctrl+B d` to detach back to dashboard
- **Kill Session** - Terminate sessions with confirmation prompt
- **Log Viewer** - View session output with scrollable viewport
- **New Session** - Create new Claude Code sessions with project path
- **Resource Monitor** - Real-time CPU/Memory usage per session
- **Filter/Search** - Filter sessions by name, project, or status
- **Auto-refresh** - Session list updates every 2 seconds
- **Config** - Customizable via `~/.claude-dashboard/config.yaml`

## Requirements

- Go 1.25+
- tmux

## Install

```bash
# From source
go install github.com/seunggabi/claude-dashboard/cmd/claude-dashboard@latest

# Or build locally
git clone https://github.com/seunggabi/claude-dashboard.git
cd claude-dashboard
make install
```

## Usage

```bash
claude-dashboard
```

## Keybindings

| Key | Action |
|-----|--------|
| `enter` | Attach to session |
| `n` | New session |
| `K` | Kill session (with confirm) |
| `l` | View logs |
| `d` | Session detail |
| `/` | Filter |
| `r` | Refresh |
| `?` | Help |
| `q` | Quit |
| `↑/k` | Move up |
| `↓/j` | Move down |
| `esc` | Go back |

## Session Naming

Sessions managed by claude-dashboard use the `cd-` prefix (e.g., `cd-my-project`).
Existing tmux sessions containing "claude" in the name are also detected.

## Config

`~/.claude-dashboard/config.yaml`:

```yaml
refresh_interval: 2s
session_prefix: "cd-"
default_dir: ""
log_history: 1000
theme: dark
```

## License

MIT

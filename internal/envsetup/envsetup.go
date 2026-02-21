package envsetup

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const marker = "<!-- claude-dashboard:agent-mail -->"

// agentMailBlock is the instruction block appended to ~/.claude/CLAUDE.md.
// It uses conditional language so agents gracefully skip it when the server is not running.
func agentMailBlock(port int) string {
	return fmt.Sprintf(`
%s
## Agent Communication (claude-dashboard)

If an MCP agent mail server is available at http://localhost:%d:
1. Register yourself on session start: call `+"`register_agent`"+` using your tmux session name as agent_id
2. Check your inbox: call `+"`fetch_inbox`"+` to see pending messages
3. Before editing any shared file, reserve it: call `+"`file_reservation_paths`"+`
4. Check inbox again when completing each major task
5. Use the beads issue ID as thread_id when messaging about a task (e.g. thread_id="claude-dashboard-88i")

If the server is not reachable, skip these steps silently.
`, marker, port)
}

// AppendAgentMailInstructions appends agent mail instructions to ~/.claude/CLAUDE.md.
// It is idempotent — if the marker is already present the file is not modified.
func AppendAgentMailInstructions(port int) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not determine home directory: %w", err)
	}

	claudeDir := filepath.Join(home, ".claude")
	if err := os.MkdirAll(claudeDir, 0755); err != nil {
		return fmt.Errorf("could not create ~/.claude: %w", err)
	}

	claudeMd := filepath.Join(claudeDir, "CLAUDE.md")

	// Read existing content
	existing := ""
	if data, err := os.ReadFile(claudeMd); err == nil {
		existing = string(data)
	}

	// Idempotency check
	if strings.Contains(existing, marker) {
		return nil
	}

	f, err := os.OpenFile(claudeMd, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open ~/.claude/CLAUDE.md: %w", err)
	}
	defer f.Close()

	_, err = f.WriteString(agentMailBlock(port))
	return err
}

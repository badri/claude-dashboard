package envsetup

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const setupFile = ".claude/worktree-setup.md"
const teardownFile = ".claude/worktree-teardown.md"

// ScaffoldCLAUDEMd writes a CLAUDE.md to sessionDir with session metadata and
// agent mail instructions. If projectDir contains a worktree-setup.md it
// references that file too.
func ScaffoldCLAUDEMd(sessionDir, sessionName, projectDir string, agentMailPort int) error {
	if err := os.MkdirAll(sessionDir, 0755); err != nil {
		return fmt.Errorf("failed to create session dir: %w", err)
	}

	content := buildContent(sessionName, projectDir, agentMailPort)

	dest := filepath.Join(sessionDir, "CLAUDE.md")
	if err := os.WriteFile(dest, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write CLAUDE.md: %w", err)
	}
	return nil
}

func buildContent(sessionName, projectDir string, agentMailPort int) string {
	content := fmt.Sprintf("# Session: %s\nCreated: %s\n", sessionName, time.Now().Format("2006-01-02"))

	if projectDir != "" {
		content += fmt.Sprintf("Project: %s\n", projectDir)
	}

	// Reference worktree setup/teardown if present
	if projectDir != "" {
		if _, err := os.Stat(filepath.Join(projectDir, setupFile)); err == nil {
			content += fmt.Sprintf(`
## Setup
Before starting work, read and execute the steps in: %s
`, setupFile)
		}
		if _, err := os.Stat(filepath.Join(projectDir, teardownFile)); err == nil {
			content += fmt.Sprintf(`
## Teardown
When done with this session, read and execute: %s
`, teardownFile)
		}
	}

	// Agent communication instructions
	content += fmt.Sprintf(`
## Agent Communication
An MCP agent mail server is running at http://localhost:%d

On session start:
1. Register yourself: call `+"`register_agent`"+` with project=<project-name> and agent_id=%q
2. Check your inbox: call `+"`fetch_inbox`"+` to see pending messages
3. Before editing any shared file, reserve it: call `+"`file_reservation_paths`"+`
4. Check inbox again when completing each major task

Use the beads issue ID as the thread_id when sending messages (e.g. thread_id="claude-dashboard-88i").
`, agentMailPort, sessionName)

	return content
}

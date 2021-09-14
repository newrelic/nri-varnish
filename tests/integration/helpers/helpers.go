//go:build integration
// +build integration

package helpers

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
	"testing"
	"time"
)

const (
	// Arbitrary amount of time to let tests exit cleanly before main process terminates.
	timeoutGracePeriod = 10 * time.Second
)

// ExecInContainer executes the given command inside the specified container. It returns three values:
// 1st - Standard Output
// 2nd - Standard Error
// 3rd - Runtime error, if any
func ExecInContainer(t *testing.T, timeout time.Duration, container string, command []string, envVars ...string) (string, string) {
	t.Helper()

	cmdLine := make([]string, 0, 3+len(command))
	cmdLine = append(cmdLine, "exec", "-i")

	for _, envVar := range envVars {
		cmdLine = append(cmdLine, "-e", envVar)
	}

	cmdLine = append(cmdLine, container)
	cmdLine = append(cmdLine, command...)

	t.Logf("executing: docker %s", strings.Join(cmdLine, " "))

	ctx, _ := context.WithTimeout(contextWithDeadline(t), timeout)
	cmd := exec.CommandContext(ctx, "docker", cmdLine...)

	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	if err := cmd.Run(); err != nil {
		t.Errorf("Integration command failed to run: %s", err)
		t.Fail()
	}

	stdout := outBuffer.String()
	stderr := errBuffer.String()

	return stdout, stderr
}

// contextWithDeadline returns context which will timeout before t.Deadline().
func contextWithDeadline(t *testing.T) context.Context {
	t.Helper()

	deadline, ok := t.Deadline()
	if !ok {
		return context.Background()
	}

	ctx, cancel := context.WithDeadline(context.Background(), deadline.Truncate(timeoutGracePeriod))

	t.Cleanup(cancel)

	return ctx
}

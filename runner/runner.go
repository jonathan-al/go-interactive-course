package runner

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Result struct {
	Output   string `json:"output"`
	Expected string `json:"expected"`
	Passed   bool   `json:"passed"`
	Error    string `json:"error,omitempty"`
}

type RunConfig struct {
	Timeout       time.Duration
	SortOutput    bool
	ContentDir    string
}

func DefaultConfig(contentDir string) RunConfig {
	return RunConfig{
		Timeout:    5 * time.Second,
		SortOutput: false,
		ContentDir: contentDir,
	}
}

func Run(code string, lessonID, exerciseID string, cfg RunConfig) Result {
	tmpDir, err := os.MkdirTemp("", "go-exercise-*")
	if err != nil {
		return Result{Error: fmt.Sprintf("failed to create temp dir: %v", err)}
	}
	defer os.RemoveAll(tmpDir)

	srcFile := filepath.Join(tmpDir, "main.go")
	if err := os.WriteFile(srcFile, []byte(code), 0644); err != nil {
		return Result{Error: fmt.Sprintf("failed to write source: %v", err)}
	}

	goModContent := "module exercise\n\ngo 1.21\n"
	goModFile := filepath.Join(tmpDir, "go.mod")
	if err := os.WriteFile(goModFile, []byte(goModContent), 0644); err != nil {
		return Result{Error: fmt.Sprintf("failed to write go.mod: %v", err)}
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "run", srcFile)
	cmd.Dir = tmpDir

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		return Result{Error: "execution timed out (5s limit)"}
	}
	if err != nil {
		errOutput := stderr.String()
		if errOutput == "" {
			errOutput = err.Error()
		}
		return Result{Error: errOutput}
	}

	actualOutput := strings.TrimRight(stdout.String(), "\n")
	expectedOutput := loadExpected(cfg.ContentDir, lessonID, exerciseID)

	if cfg.SortOutput {
		actualOutput = sortLines(actualOutput)
		expectedOutput = sortLines(expectedOutput)
	}

	passed := actualOutput == expectedOutput

	return Result{
		Output:   actualOutput,
		Expected: expectedOutput,
		Passed:   passed,
	}
}

func loadExpected(contentDir, lessonID, exerciseID string) string {
	path := filepath.Join(contentDir, "lessons", lessonID, "exercises", exerciseID, "expected.txt")
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.TrimRight(string(data), "\n")
}

func sortLines(s string) string {
	lines := strings.Split(s, "\n")
	sort.Strings(lines)
	return strings.Join(lines, "\n")
}

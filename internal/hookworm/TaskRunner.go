package hookworm

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// HookExecutionError represents an error from a failed hook execution
type HookExecutionError struct {
	HookName string
	ExitCode int
}

func (e *HookExecutionError) Error() string {
	return fmt.Sprintf("hook %s failed with exit code %d", e.HookName, e.ExitCode)
}

// ExecuteTasks loads and executes hooks from .hookworm.yaml
func ExecuteTasks(config *TaskBook) error {

	for i, hook := range config.Task {
		log.Printf("Task %d: Name=%s, Command=%s", i+1, hook.Name, hook.Command)

		// Execute the command
		cmd := exec.Command("sh", "-c", hook.Command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		// Check the exit code
		var exitCode int
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				exitCode = exitErr.ExitCode()
			} else {
				return fmt.Errorf("executing hook %s: %v", hook.Name, err)
			}
		}

		log.Printf("Task %s completed with exit code %d", hook.Name, exitCode)

		// Fail fast if exit code is non-zero
		if exitCode != 0 {
			return &HookExecutionError{
				HookName: hook.Name,
				ExitCode: exitCode,
			}
		}
	}

	return nil
}

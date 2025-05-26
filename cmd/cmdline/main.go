package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tastybug/hookworm/internal/hookworm"
)

func main() {

	// Define flags
	configPtr := flag.String("config", ".hookworm.yml", "Path to the task book config file")
	flag.Usage = printUsage
	flag.Parse()

	// Check for a command
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	command := args[0]

	printWorkingDirectory()

	switch command {
	case "install":
		if err := hookworm.InstallHook(); err != nil {
			log.Printf("Error installing hook: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Hookworm installed successfully")
	case "advise":
		taskBook := loadTaskBookOrFail(configPtr)
		fmt.Println("Log checks only..")
		if err := hookworm.ExecuteTasks(taskBook, false); err != nil {
			log.Printf("Error running hooks: %v\n", err)
			os.Exit(1)
		}
	case "assert":
		taskBook := loadTaskBookOrFail(configPtr)
		fmt.Println("Log and assert checks..")
		if err := hookworm.ExecuteTasks(taskBook, true); err != nil {
			log.Printf("Error running hooks: %v\n", err)
			os.Exit(1)
		}
	case "help":
		printUsage()
		os.Exit(0)
	default:
		printUsage()
		os.Exit(1)
	}
}

func loadTaskBookOrFail(configPtr *string) *hookworm.TaskBook {
	taskBook, err := hookworm.InitializeTaskBook(*configPtr)
	if err != nil {
		log.Printf("Error loading task book: %v\n", err)
		os.Exit(1)
	}
	return taskBook
}

func printWorkingDirectory() {
	cwd, err := os.Getwd()
	if err != nil {
		_ = fmt.Errorf("error getting current working directory: %v", err)
		os.Exit(1)
	}
	log.Printf("Current working directory: %s", cwd)
}

// printUsage displays the command-line usage
func printUsage() {
	_, _ = fmt.Printf("Usage: hookworm [flags] <command>\n")
	_, _ = fmt.Printf("Commands:\n")
	_, _ = fmt.Printf("  install  Install hookworm as an assertive Git pre-commit hook. Must be run in a working copy folder.\n")
	_, _ = fmt.Printf("  advise   Run checks, log findings but does not fail.\n")
	_, _ = fmt.Printf("  assert   Run checks, log findings and fail on issues.\n")
	_, _ = fmt.Printf("  help     Show this help.\n")
	_, _ = fmt.Printf("Flags:\n")
	flag.PrintDefaults()
}

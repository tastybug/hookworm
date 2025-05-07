package main

import (
	"flag"
	"fmt"
	"github.com/tastybug/hookworm/internal/hookworm"
	"log"
	"os"
)

func main() {

	// Define flags
	configPtr := flag.String("config", ".hookworm.yml", "Path to the hookplay config file")
	flag.Usage = printUsage
	flag.Parse()

	// Check for a command
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	command := args[0]

	err := printWorkingDirectory()
	taskBook, err := hookworm.InitializeTaskBook(*configPtr)
	if err != nil {
		fmt.Errorf("Error loading taskBook: %v", err)
		os.Exit(1)
	}

	switch command {
	case "install":
		if err := hookworm.InstallHook(); err != nil {
			fmt.Fprintf(os.Stderr, "Error installing hook: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Hookworm installed successfully")
	case "test":
		fmt.Println("Dry-run..")
		if err := hookworm.ExecuteTasks(taskBook); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error running hooks: %v\n", err)
			os.Exit(1)
		}
	case "trigger":
		fmt.Println("Triggered....")
		if err := hookworm.ExecuteTasks(taskBook); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error running hooks: %v\n", err)
			os.Exit(1)
		}
	default:
		printUsage()
		os.Exit(1)
	}
}

func printWorkingDirectory() error {
	cwd, err := os.Getwd()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error getting current working directory: %v\n", err)
		os.Exit(1)
	}
	log.Printf("Current working directory: %s", cwd)
	return err
}

// printUsage displays the command-line usage
func printUsage() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage: hookworm [flags] <command>\n")
	_, _ = fmt.Fprintf(os.Stderr, "Commands:\n")
	_, _ = fmt.Fprintf(os.Stderr, "  install  Install hookworm as a Git pre-commit hook\n")
	_, _ = fmt.Fprintf(os.Stderr, "  test     Manually test the hookworm checks\n")
	_, _ = fmt.Fprintf(os.Stderr, "  trigger  Run the hookworm checks (called by Git pre-commit hook)\n")
	_, _ = fmt.Fprintf(os.Stderr, "Flags:\n")
	_, _ = fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

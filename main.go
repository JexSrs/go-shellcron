package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/robfig/cron/v3"
)

const (
	DefaultScriptsDir = "./scripts"
)

func main() {
	scriptsDir := getScriptsDir()

	c := cron.New()

	files, err := os.ReadDir(scriptsDir)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".sh") {
			log.Printf("Skipping non-shell file: %s", file.Name())
			continue
		}

		filePath := fmt.Sprintf("%s/%s", scriptsDir, file.Name())
		cronLine, err := parseCronLine(filePath)
		if err != nil {
			log.Printf("Error parsing cron line from %s: %v", file.Name(), err)
			continue
		}

		_, err = c.AddFunc(cronLine, func() { runScript(filePath) })
		if err != nil {
			log.Printf("Error scheduling script %s: %v", file.Name(), err)
			continue
		}
		log.Printf("Scheduled %s with cron spec %s", file.Name(), cronLine)
	}

	c.Start()
	defer c.Stop()

	// Prevent the program from exiting.
	select {}
}

func getScriptsDir() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return DefaultScriptsDir
}

func parseCronLine(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cronRegex := regexp.MustCompile(`^#CRON: (.*)$`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := cronRegex.FindStringSubmatch(line)
		if len(matches) > 1 {
			return matches[1], nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("no cron line found in file: %s", filePath)
}

func runScript(scriptPath string) {
	cmd := exec.Command("/bin/sh", scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Printf("Error running script %s: %s", scriptPath, err)
		return
	}
	log.Printf("Finished running script: %s", scriptPath)
}

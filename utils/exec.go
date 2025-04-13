package utils

import (
	"io"
	"log"
	"os/exec"
)

func Run(cmd string, args ...string) {
	command := exec.Command(cmd, args...)

	if !CommandExists(cmd) {
		log.Fatalf("The commande %s does not exist", cmd)
		return
	}
	stdout, err := command.StdoutPipe()
	if err != nil {
		log.Fatalf("Error getting stdout: %v", err)
	}

	stderr, err := command.StderrPipe()
	if err != nil {
		log.Fatalf("Error getting stderr: %v", err)
	}

	if err := command.Start(); err != nil {
		log.Fatalf("Command start failed: %v", err)
	}

	// Copy output to terminal as it happens
	go io.Copy(log.Writer(), stdout)
	go io.Copy(log.Writer(), stderr)

	if err := command.Wait(); err != nil {
		log.Fatalf("Command failed: %v", err)
	}
}

func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// Check if a given package is already installed on the system
func IsInstalled(pkg string) bool {
	return CommandExists(pkg)
}

package myutils

import (
	"os"
	"os/exec"
)

// ClearScreen ...
// Clears the terminal.
func ClearScreen() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}

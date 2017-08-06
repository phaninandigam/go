package myutils

import (
	"os"
	"os/exec"
)

func ClearScreen() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	g := NewGame(20, 60)

	for i := 0; i < 100; i++ {
		clearTerm()
		g.Step()
		fmt.Print(g)
		time.Sleep(time.Millisecond * 150)
	}
}

func clearTerm() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return
	}
	fmt.Printf("\x1b[2J")
}

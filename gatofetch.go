package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func displaySystemInfo() {
	fmt.Println("\033[1mKernel:\033[0m", runCommand("uname -r"))
	fmt.Println("\033[1mUser:\033[0m", runCommand("whoami"))
	fmt.Println("\033[1mDE:\033[0m", getDesktopEnvironment())
	fmt.Println("\033[1mCPU:\033[0m", runCommand("grep 'model name' /proc/cpuinfo | uniq | awk -F: '{print $2}'"))
	fmt.Printf("\033[1mRAM:\033[0m %s\n", getRAMInfo())
}

func displayKittyFace() {
	fmt.Println("  /\\_/\\")
	fmt.Println(" ( o.o )")
	fmt.Println("  > ^ <\n")
}

func getDesktopEnvironment() string {
	return os.Getenv("XDG_CURRENT_DESKTOP")
}

func getRAMInfo() string {
	total := runCommand("free -m | awk '/Mem:/ {print $2}'")
	used := runCommand("free -m | awk '/Mem:/ {print $3}'")
	return fmt.Sprintf("%s MB / %s MB", used, total)
}

func runCommand(command string) string {
	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return strings.TrimSpace(string(output))
}

func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	displaySystemInfo()
	displayKittyFace()
}


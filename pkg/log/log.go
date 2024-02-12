package log

import (
	"fmt"
	"strings"
)

func Info(message string) {
	fmt.Printf("INFO: %s\n", message)
}

func Error(message string) {
	fmt.Printf("ERROR: %s\n", message)
}

func CMD(command ...string) {
	fmt.Printf("[CMD]: %s\n", strings.Join(command, " "))
}

func Usage() {
	fmt.Println("Usage: gollo <SUBCOMMAND> [ARGS] ./examples/foo.glo")
	fmt.Println("SUBCOMMANDS:")
	fmt.Println("    run        Simulately run program")
	fmt.Println("    compile    Compile program into an executable")
	fmt.Println("ARGS:")
	fmt.Println("    -r         Instantly run compiled program after compilation")
}

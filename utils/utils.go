package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// AskYesNo prompts the user with a question and returns true if the answer is yes
func AskYesNo(question string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [y/N]: ", question)
	answer, _ := reader.ReadString('\n')
	answer = strings.ToLower(strings.TrimSpace(answer))
	return answer == "y" || answer == "yes"
}

func AskForString(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", question)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

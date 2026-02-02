package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter 0x.. / 0b.. / 0o.. (no prefix = hex) to convert to decimal, or 'stop' to exit:")

	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		//EOF or error
		if !input.Scan() {
			break
		}
		line := strings.TrimSpace(input.Text())
		if line == "" {
			continue
		}

		if strings.EqualFold(line, "stop") {
			return
		}

		digits, base, ok := parseWithPrefix(line)
		if !ok {
			fmt.Println("failed!")
			continue
		}

		var n big.Int

		if _, ok := n.SetString(digits, base); !ok {
			fmt.Println("failed!")
			continue
		}
		fmt.Println(&n)
	}

	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "input error:", err)
	}
}

func parseWithPrefix(s string) (digits string, base int, ok bool) {
	s = strings.TrimSpace(s)
	if len(s) < 2 {
		return s, 16, true
	}

	low := strings.ToLower(s[:2])
	switch low {
	case "0x":
		return s[2:], 16, len(s) > 2
	case "0b":
		return s[2:], 2, len(s) > 2
	case "0o":
		return s[2:], 8, len(s) > 2
	default:
		return s, 16, true
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Select ~> ")
		fmt.Println("1) S = V * t")
		fmt.Println("2) ρ = m / t")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		sel, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error!")
			return
		}

		switch sel {
		case 1:
			speed()
		case 2:
			density()
		}
	}
}

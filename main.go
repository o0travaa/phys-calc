package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("~> ")

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
		}
	}
}

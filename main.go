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
		fmt.Println("1) S = Vt")
		fmt.Println("2) ρ = m / V")
		fmt.Println("3) F = mg")
		fmt.Println("4) Fупр = kΔx")
		fmt.Println("5) p = F / S")
		fmt.Println("6) p = ρgh")
		fmt.Println("7) Fа = ρgV")
		fmt.Println("8) A = FS")
		fmt.Println("9) N = A / t")
		fmt.Println("10) η = Aп / Aз 100%")
		fmt.Println("11) Q = cm(t2 - t1)")
		fmt.Println("12) Q = λm")
		fmt.Print("Select ~> ")

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
		case 3:
			gravity()
		case 4:
			elasticForce()
		case 5:
			pressure()
		case 6:
			hydrostaticPressure()
		case 7:
			archimedes()
		case 8:
			work()
		case 9:
			power()
		case 10:
			efficiency()
		case 11:
			heatCapacity()
		case 12:
			heatFusion()
		default:
			fmt.Println("Error!")
		}
	}
}

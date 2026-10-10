package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func speed() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter \"t\" (or press Enter if empty) ~> ")
	speed_t, _ := reader.ReadString('\n')
	speed_t = strings.TrimSpace(speed_t)

   	fmt.Print("Enter \"V\" (or press Enter if empty) ~> ")
	speed_V, _ := reader.ReadString('\n')
	speed_V = strings.TrimSpace(speed_V)

   	fmt.Print("Enter \"S\" (or press Enter if empty) ~> ")
	speed_S, _ := reader.ReadString('\n')
	speed_S = strings.TrimSpace(speed_S)

	if speed_S == "" && speed_t != "" && speed_V != "" {
		t, err1 := strconv.Atoi(speed_t)
		V, err2 := strconv.Atoi(speed_V)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: S = %d\n", t * V)
	} else if speed_V == "" && speed_t != "" && speed_S != "" {
		t, err1 := strconv.Atoi(speed_t)
		S, err2 := strconv.Atoi(speed_S)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: V = %d\n", S / t)
	} else if speed_t == "" && speed_V != "" && speed_S != "" {
		V, err1 := strconv.Atoi(speed_V)
		S, err2 := strconv.Atoi(speed_S)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: t = %d\n", S / V)
	} else {
		fmt.Println("Error!")
	}
}

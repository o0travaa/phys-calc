package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		fmt.Printf("Result: S = %d\n", t*V)
	} else if speed_V == "" && speed_t != "" && speed_S != "" {
		t, err1 := strconv.Atoi(speed_t)
		S, err2 := strconv.Atoi(speed_S)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: V = %d\n", S/t)
	} else if speed_t == "" && speed_V != "" && speed_S != "" {
		V, err1 := strconv.Atoi(speed_V)
		S, err2 := strconv.Atoi(speed_S)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: t = %d\n", S/V)
	} else {
		fmt.Println("Error!")
	}
}

func density() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter \"ρ\" (or press Enter if empty) ~> ")
	density_p, _ := reader.ReadString('\n')
	density_p = strings.TrimSpace(density_p)

	fmt.Print("Enter \"V\" (or press Enter if empty) ~> ")
	density_V, _ := reader.ReadString('\n')
	density_V = strings.TrimSpace(density_V)

	fmt.Print("Enter \"m\" (or press Enter if empty) ~> ")
	density_m, _ := reader.ReadString('\n')
	density_m = strings.TrimSpace(density_m)

	if density_p == "" && density_V != "" && density_m != "" {
		m, err1 := strconv.Atoi(density_m)
		V, err2 := strconv.Atoi(density_V)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: ρ = %d\n", m/V)
	} else if density_V == "" && density_p != "" && density_m != "" {
		p, err1 := strconv.Atoi(density_p)
		m, err2 := strconv.Atoi(density_m)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: V = %d\n", m/p)
	} else if density_m == "" && density_p != "" && density_V != "" {
		p, err1 := strconv.Atoi(density_p)
		V, err2 := strconv.Atoi(density_V)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: m = %d\n", p/V)
	} else {
		fmt.Println("Error!")
	}
}

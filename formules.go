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
	speedT, _ := reader.ReadString('\n')
	speedT = strings.TrimSpace(speedT)

	fmt.Print("Enter \"V\" (or press Enter if empty) ~> ")
	speedV, _ := reader.ReadString('\n')
	speedV = strings.TrimSpace(speedV)

	fmt.Print("Enter \"S\" (or press Enter if empty) ~> ")
	speedS, _ := reader.ReadString('\n')
	speedS = strings.TrimSpace(speedS)

	if speedS == "" && speedT != "" && speedV != "" {
		t, err1 := strconv.ParseFloat(speedT, 64)
		V, err2 := strconv.ParseFloat(speedV, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: S = %.3f\n", t*V)
	} else if speedV == "" && speedT != "" && speedS != "" {
		t, err1 := strconv.ParseFloat(speedT, 64)
		S, err2 := strconv.ParseFloat(speedS, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: V = %.3f\n", S/t)
	} else if speedT == "" && speedV != "" && speedS != "" {
		V, err1 := strconv.ParseFloat(speedV, 64)
		S, err2 := strconv.ParseFloat(speedS, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: t = %.3f\n", S/V)
	} else {
		fmt.Println("Error!")
	}
}

func density() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter \"ρ\" (or press Enter if empty) ~> ")
	densityP, _ := reader.ReadString('\n')
	densityP = strings.TrimSpace(densityP)

	fmt.Print("Enter \"V\" (or press Enter if empty) ~> ")
	densityV, _ := reader.ReadString('\n')
	densityV = strings.TrimSpace(densityV)

	fmt.Print("Enter \"m\" (or press Enter if empty) ~> ")
	densityM, _ := reader.ReadString('\n')
	densityM = strings.TrimSpace(densityM)

	if densityP == "" && densityV != "" && densityM != "" {
		m, err1 := strconv.ParseFloat(densityM, 64)
		V, err2 := strconv.ParseFloat(densityV, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: ρ = %.3f\n", m/V)
	} else if densityV == "" && densityP != "" && densityM != "" {
		p, err1 := strconv.ParseFloat(densityP, 64)
		m, err2 := strconv.ParseFloat(densityM, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: V = %.3f\n", m/p)
	} else if densityM == "" && densityP != "" && densityV != "" {
		p, err1 := strconv.ParseFloat(densityP, 64)
		V, err2 := strconv.ParseFloat(densityV, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: m = %.3f\n", p/V)
	} else {
		fmt.Println("Error!")
	}
}

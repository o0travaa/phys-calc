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
		fmt.Printf("Result: m = %.3f\n", p*V)
	} else {
		fmt.Println("Error!")
	}
}

func gravity() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter \"F\" (or press Enter if empty) ~> ")
	gravityF, _ := reader.ReadString('\n')
	gravityF = strings.TrimSpace(gravityF)

	fmt.Print("Enter \"m\" (or press Enter if empty) ~> ")
	gravityM, _ := reader.ReadString('\n')
	gravityM = strings.TrimSpace(gravityM)

	if gravityF == "" && gravityM != "" {
		m, err := strconv.ParseFloat(gravityM, 64)
		if err != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: F = %.3f\n", m*10)
	} else if gravityM == "" && gravityF != "" {
		F, err := strconv.ParseFloat(gravityF, 64)
		if err != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: m = %.3f\n", F/10)
	} else {
		fmt.Println("Error!")
	}
}

func elasticForce() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter \"F\" (or press Enter if empty) ~> ")
	elasticF, _ := reader.ReadString('\n')
	elasticF = strings.TrimSpace(elasticF)

	fmt.Print("Enter \"k\" (or press Enter if empty) ~> ")
	elasticK, _ := reader.ReadString('\n')
	elasticK = strings.TrimSpace(elasticK)

	fmt.Print("Enter \"Δl\" (or press Enter if empty) ~> ")
	elasticL, _ := reader.ReadString('\n')
	elasticL = strings.TrimSpace(elasticL)

	if elasticF == "" && elasticK != "" && elasticL != "" {
		k, err1 := strconv.ParseFloat(elasticK, 64)
		l, err2 := strconv.ParseFloat(elasticL, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: F = %.3f\n", k*l)
	} else if elasticK == "" && elasticF != "" && elasticL != "" {
		F, err1 := strconv.ParseFloat(elasticF, 64)
		l, err2 := strconv.ParseFloat(elasticL, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: k = %.3f\n", F/l)
	} else if elasticL == "" && elasticF != "" && elasticK != "" {
		F, err1 := strconv.ParseFloat(elasticF, 64)
		k, err2 := strconv.ParseFloat(elasticK, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: Δl = %.3f\n", F/k)
	} else {
		fmt.Println("Error!")
	}
}

func pressure() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter \"p\" (or press Enter if empty) ~> ")
	pressureP, _ := reader.ReadString('\n')
	pressureP = strings.TrimSpace(pressureP)

	fmt.Print("Enter \"F\" (or press Enter if empty) ~> ")
	pressureF, _ := reader.ReadString('\n')
	pressureF = strings.TrimSpace(pressureF)

	fmt.Print("Enter \"S\" (or press Enter if empty) ~> ")
	pressureS, _ := reader.ReadString('\n')
	pressureS = strings.TrimSpace(pressureS)

	if pressureP == "" && pressureF != "" && pressureS != "" {
		F, err1 := strconv.ParseFloat(pressureF, 64)
		S, err2 := strconv.ParseFloat(pressureS, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: p = %.3f\n", F/S)
	} else if pressureF == "" && pressureP != "" && pressureS != "" {
		p, err1 := strconv.ParseFloat(pressureP, 64)
		S, err2 := strconv.ParseFloat(pressureS, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: F = %.3f\n", p*S)
	} else if pressureS == "" && pressureP != "" && pressureF != "" {
		p, err1 := strconv.ParseFloat(pressureP, 64)
		F, err2 := strconv.ParseFloat(pressureF, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: S = %.3f\n", F/p)
	} else {
		fmt.Println("Error!")
	}
}

func hydrostaticPressure() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter \"p\" (or press Enter if empty) ~> ")
	hydroP, _ := reader.ReadString('\n')
	hydroP = strings.TrimSpace(hydroP)

	fmt.Print("Enter \"ρ\" (or press Enter if empty) ~> ")
	hydroRho, _ := reader.ReadString('\n')
	hydroRho = strings.TrimSpace(hydroRho)

	fmt.Print("Enter \"h\" (or press Enter if empty) ~> ")
	hydroH, _ := reader.ReadString('\n')
	hydroH = strings.TrimSpace(hydroH)

	if hydroP == "" && hydroRho != "" && hydroH != "" {
		rho, err1 := strconv.ParseFloat(hydroRho, 64)
		h, err2 := strconv.ParseFloat(hydroH, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: p = %.3f\n", rho*10*h)
	} else if hydroRho == "" && hydroP != "" && hydroH != "" {
		p, err1 := strconv.ParseFloat(hydroP, 64)
		h, err2 := strconv.ParseFloat(hydroH, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: ρ = %.3f\n", p/(10*h))
	} else if hydroH == "" && hydroP != "" && hydroRho != "" {
		p, err1 := strconv.ParseFloat(hydroP, 64)
		rho, err2 := strconv.ParseFloat(hydroRho, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: h = %.3f\n", p/(rho*10))
	} else {
		fmt.Println("Error!")
	}
}

func archimedes() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter \"F\" (or press Enter if empty) ~> ")
	archF, _ := reader.ReadString('\n')
	archF = strings.TrimSpace(archF)

	fmt.Print("Enter \"ρ\" (or press Enter if empty) ~> ")
	archRho, _ := reader.ReadString('\n')
	archRho = strings.TrimSpace(archRho)

	fmt.Print("Enter \"V\" (or press Enter if empty) ~> ")
	archV, _ := reader.ReadString('\n')
	archV = strings.TrimSpace(archV)

	if archF == "" && archRho != "" && archV != "" {
		rho, err1 := strconv.ParseFloat(archRho, 64)
		V, err2 := strconv.ParseFloat(archV, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: F = %.3f\n", rho*10*V)
	} else if archRho == "" && archF != "" && archV != "" {
		F, err1 := strconv.ParseFloat(archF, 64)
		V, err2 := strconv.ParseFloat(archV, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: ρ = %.3f\n", F/(10*V))
	} else if archV == "" && archF != "" && archRho != "" {
		F, err1 := strconv.ParseFloat(archF, 64)
		rho, err2 := strconv.ParseFloat(archRho, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: V = %.3f\n", F/(rho*10))
	} else {
		fmt.Println("Error!")
	}
}

func work() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter \"A\" (or press Enter if empty) ~> ")
	workA, _ := reader.ReadString('\n')
	workA = strings.TrimSpace(workA)

	fmt.Print("Enter \"F\" (or press Enter if empty) ~> ")
	workF, _ := reader.ReadString('\n')
	workF = strings.TrimSpace(workF)

	fmt.Print("Enter \"S\" (or press Enter if empty) ~> ")
	workS, _ := reader.ReadString('\n')
	workS = strings.TrimSpace(workS)

	if workA == "" && workF != "" && workS != "" {
		F, err1 := strconv.ParseFloat(workF, 64)
		S, err2 := strconv.ParseFloat(workS, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: A = %.3f\n", F*S)
	} else if workF == "" && workA != "" && workS != "" {
		A, err1 := strconv.ParseFloat(workA, 64)
		S, err2 := strconv.ParseFloat(workS, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: F = %.3f\n", A/S)
	} else if workS == "" && workA != "" && workF != "" {
		A, err1 := strconv.ParseFloat(workA, 64)
		F, err2 := strconv.ParseFloat(workF, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: S = %.3f\n", A/F)
	} else {
		fmt.Println("Error!")
	}
}

func power() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter \"N\" (or press Enter if empty) ~> ")
	powerN, _ := reader.ReadString('\n')
	powerN = strings.TrimSpace(powerN)

	fmt.Print("Enter \"A\" (or press Enter if empty) ~> ")
	powerA, _ := reader.ReadString('\n')
	powerA = strings.TrimSpace(powerA)

	fmt.Print("Enter \"t\" (or press Enter if empty) ~> ")
	powerT, _ := reader.ReadString('\n')
	powerT = strings.TrimSpace(powerT)

	if powerN == "" && powerA != "" && powerT != "" {
		A, err1 := strconv.ParseFloat(powerA, 64)
		t, err2 := strconv.ParseFloat(powerT, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: N = %.3f\n", A/t)
	} else if powerA == "" && powerN != "" && powerT != "" {
		N, err1 := strconv.ParseFloat(powerN, 64)
		t, err2 := strconv.ParseFloat(powerT, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: A = %.3f\n", N*t)
	} else if powerT == "" && powerN != "" && powerA != "" {
		N, err1 := strconv.ParseFloat(powerN, 64)
		A, err2 := strconv.ParseFloat(powerA, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: t = %.3f\n", A/N)
	} else {
		fmt.Println("Error!")
	}
}

func efficiency() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter \"η\" (or press Enter if empty) ~> ")
	effEta, _ := reader.ReadString('\n')
	effEta = strings.TrimSpace(effEta)

	fmt.Print("Enter \"Aп\" (or press Enter if empty) ~> ")
	effUseful, _ := reader.ReadString('\n')
	effUseful = strings.TrimSpace(effUseful)

	fmt.Print("Enter \"Aз\" (or press Enter if empty) ~> ")
	effSpent, _ := reader.ReadString('\n')
	effSpent = strings.TrimSpace(effSpent)

	if effEta == "" && effUseful != "" && effSpent != "" {
		useful, err1 := strconv.ParseFloat(effUseful, 64)
		spent, err2 := strconv.ParseFloat(effSpent, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: η = %.3f\n", useful/spent*100)
	} else if effUseful == "" && effEta != "" && effSpent != "" {
		eta, err1 := strconv.ParseFloat(effEta, 64)
		spent, err2 := strconv.ParseFloat(effSpent, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: Aп = %.3f\n", eta/100*spent)
	} else if effSpent == "" && effEta != "" && effUseful != "" {
		eta, err1 := strconv.ParseFloat(effEta, 64)
		useful, err2 := strconv.ParseFloat(effUseful, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: Aз = %.3f\n", useful/(eta/100))
	} else {
		fmt.Println("Error!")
	}
}

func heatCapacity() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter \"Q\" (or press Enter if empty) ~> ")
	heatQ, _ := reader.ReadString('\n')
	heatQ = strings.TrimSpace(heatQ)

	fmt.Print("Enter \"c\" (or press Enter if empty) ~> ")
	heatC, _ := reader.ReadString('\n')
	heatC = strings.TrimSpace(heatC)

	fmt.Print("Enter \"m\" (or press Enter if empty) ~> ")
	heatM, _ := reader.ReadString('\n')
	heatM = strings.TrimSpace(heatM)

	fmt.Print("Enter \"t1\" (or press Enter if empty) ~> ")
	heatT1, _ := reader.ReadString('\n')
	heatT1 = strings.TrimSpace(heatT1)

	fmt.Print("Enter \"t2\" (or press Enter if empty) ~> ")
	heatT2, _ := reader.ReadString('\n')
	heatT2 = strings.TrimSpace(heatT2)

	if heatQ == "" && heatC != "" && heatM != "" && heatT1 != "" && heatT2 != "" {
		c, err1 := strconv.ParseFloat(heatC, 64)
		m, err2 := strconv.ParseFloat(heatM, 64)
		t1, err3 := strconv.ParseFloat(heatT1, 64)
		t2, err4 := strconv.ParseFloat(heatT2, 64)
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: Q = %.3f\n", c*m*(t2-t1))
	} else if heatC == "" && heatQ != "" && heatM != "" && heatT1 != "" && heatT2 != "" {
		Q, err1 := strconv.ParseFloat(heatQ, 64)
		m, err2 := strconv.ParseFloat(heatM, 64)
		t1, err3 := strconv.ParseFloat(heatT1, 64)
		t2, err4 := strconv.ParseFloat(heatT2, 64)
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: c = %.3f\n", Q/(m*(t2-t1)))
	} else if heatM == "" && heatQ != "" && heatC != "" && heatT1 != "" && heatT2 != "" {
		Q, err1 := strconv.ParseFloat(heatQ, 64)
		c, err2 := strconv.ParseFloat(heatC, 64)
		t1, err3 := strconv.ParseFloat(heatT1, 64)
		t2, err4 := strconv.ParseFloat(heatT2, 64)
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: m = %.3f\n", Q/(c*(t2-t1)))
	} else if heatT1 == "" && heatQ != "" && heatC != "" && heatM != "" && heatT2 != "" {
		Q, err1 := strconv.ParseFloat(heatQ, 64)
		c, err2 := strconv.ParseFloat(heatC, 64)
		m, err3 := strconv.ParseFloat(heatM, 64)
		t2, err4 := strconv.ParseFloat(heatT2, 64)
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: t1 = %.3f\n", t2-Q/(c*m))
	} else if heatT2 == "" && heatQ != "" && heatC != "" && heatM != "" && heatT1 != "" {
		Q, err1 := strconv.ParseFloat(heatQ, 64)
		c, err2 := strconv.ParseFloat(heatC, 64)
		m, err3 := strconv.ParseFloat(heatM, 64)
		t1, err4 := strconv.ParseFloat(heatT1, 64)
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: t2 = %.3f\n", t1+Q/(c*m))
	} else {
		fmt.Println("Error!")
	}
}

func heatFusion() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter \"Q\" (or press Enter if empty) ~> ")
	fusionQ, _ := reader.ReadString('\n')
	fusionQ = strings.TrimSpace(fusionQ)

	fmt.Print("Enter \"λ\" (or press Enter if empty) ~> ")
	fusionL, _ := reader.ReadString('\n')
	fusionL = strings.TrimSpace(fusionL)

	fmt.Print("Enter \"m\" (or press Enter if empty) ~> ")
	fusionM, _ := reader.ReadString('\n')
	fusionM = strings.TrimSpace(fusionM)

	if fusionQ == "" && fusionL != "" && fusionM != "" {
		l, err1 := strconv.ParseFloat(fusionL, 64)
		m, err2 := strconv.ParseFloat(fusionM, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: Q = %.3f\n", l*m)
	} else if fusionL == "" && fusionQ != "" && fusionM != "" {
		Q, err1 := strconv.ParseFloat(fusionQ, 64)
		m, err2 := strconv.ParseFloat(fusionM, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: λ = %.3f\n", Q/m)
	} else if fusionM == "" && fusionQ != "" && fusionL != "" {
		Q, err1 := strconv.ParseFloat(fusionQ, 64)
		l, err2 := strconv.ParseFloat(fusionL, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error!")
			return
		}
		fmt.Printf("Result: m = %.3f\n", Q/l)
	} else {
		fmt.Println("Error!")
	}
}

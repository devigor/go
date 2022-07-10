package main

import "fmt"

type ContaCorrent struct {
	titular     string
	agencia     int
	numeroConta int
	saldo       float64
}

func main() {
    contaIgor := ContaCorrent{titular: "Igor", agencia: 589, numeroConta: 123456, saldo: 125.50}
    fmt.Println(contaIgor)
}

/*
	Scrivere un programma Go lez2 pari dispari.go che legge un intero n e, a
	seconda del valore di n, stampa uno dei messaggi
	n è pari
	oppure
	n è dispari
*/

package main

import (
	"fmt"
)

func pariDispari(n int) {
	if n%2 == 0 {
		fmt.Println(n, " è pari")
	} else {
		fmt.Println(n, " è dispari")
	}
}

func input() {
	for {
		fmt.Println("Inserisci un numero: ")
		var i int
		_, err := fmt.Scanf("%d", &i)
		fmt.Println("Numero inserito: ", i)
		if err != nil {
			fmt.Println("invalid input")
			break
		}
		pariDispari(i)
	}
}

func main() {
	//funzione input
	input()
}

/*
  Scrivere un programma Go lez2 maggiore.go che legge due interi e stampa il
  maggiore.
  Annotazioni In Go è possibile lo scambio diretto di valori tra due variabili:
  a, b = b, a
*/

package main

import (
	"fmt"
)

func maggiore(a, b int)  {
  if a>b {
    fmt.Println("Il maggiore è:  ", a)
  } else {
    fmt.Println("Il maggiore è:  ", b)
  }

}

func input() {
  //primo numero
	fmt.Println("Inserisci il primo numero: ")
	var first int
	fmt.Scanf("%d", &first)
  //secondo numero
	fmt.Println("Inserisci il secondo numero:  ")
	var second int
	fmt.Scanf("%d", &second)
  maggiore(first, second)
}

func main()  {
  input()
}

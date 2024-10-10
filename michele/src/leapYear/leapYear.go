/*
Funzionamento degli anni bisestili, secondo Wikipedia:
	"Premettendo che un anno si dice secolare quando è divisibile per 100, nel calendario gregoriano sono quindi bisestili:
    	- gli anni non secolari il cui numero è divisibile per 4;
    	- gli anni secolari il cui numero è divisibile per 400."
Inoltre:
	"Il calendario gregoriano fu introdotto nel 1582. Benché, in via teorica, sia possibile estenderlo anche agli anni precedenti, per questi, di norma, si usa il calendario giuliano. Perciò sono bisestili tutti gli anni divisibili per 4, compresi quelli secolari, dal 4 al 1580 dell'era volgare."	

Vuol dire che deve sussistere una di queste condizioni:
	- Divisibile per 400
	- Divisibile per 4 AND non divisibile per 100
	- Divisibile per 4 AND inferiore a 1582

*/

package main
import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Ottieni input
	var year int
	if len(os.Args) > 1 {
		year, _ = strconv.Atoi(os.Args[1])
	} else {
		fmt.Print("Inserisci l'anno: ")
		fmt.Scan(&year)
	}

	// Verifica che l'input sia della forma corretta
	if year < 0 {
		fmt.Println("Il numero che hai inserito è negativo. Se stavi cercando di inserire un anno del periodo a.C., sappi che il calendario giuliano (su cui si basa la definizione di anno bisestile fino alla sua sostituzione con quello gregoriano nel 1582) va solo fino all'anno zero e, pertanto, un anno precedente allo zero non può essere bisestile.")
	} else if year == 0 {
		fmt.Println("L'anno zero non è considerato bisestile. Il primo anno bisestile è il 4 d.C.")
	return
	} else {
		// Calcoli
		var isLeapYear bool
		isLeapYear = (year%400 == 0 || year%4 == 0 && year < 1582 || year%4 == 0 && year%100 != 0)

		// Restituisci output
		fmt.Print("L'anno che hai inserito, il ", year, " d.C., ")
		if isLeapYear {
			fmt.Print("è")
		} else {
			fmt.Print("non è")
		}
		fmt.Print(" bisestile.")

		// Aggiungi dei trivia
		if year%100 == 0 {
			fmt.Print(" Secondo il calendario gregoriano, gli anni secolari (cioè quelli divisibili per 100) non sono bisestili, neppure se sono divisibili per 4.")
			if year%400 == 0 {
				fmt.Print(" Tuttavia, fanno eccezione gli anni divisibili per 400, che sono sempre bisestili!")
			}
			if year < 1582 {
				fmt.Print(" Tuttavia, fino al 1582 si usò il calendario giuliano, che non segue questa regola!")
			}
		}
	}
	fmt.Print("\n")
}
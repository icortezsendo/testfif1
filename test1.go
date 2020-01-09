package main

import (
	"fmt"
	"strconv"
)

const larTipo int = 3
const larLargo int = 5

func main() {
	mapa, err := separaDatos("A0511AB398765UJ1N230200")
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println(mapa)
	}
}

func separaDatos(datoEntrada string) (map[string]string, string) {

	mapa := make(map[string]string)
	var errores string = ""

	if datoEntrada == "" {
		errores = "No se introdujo el dato a leer"
		return mapa, errores
	}

	var largoTotal int = len(datoEntrada)
	//fmt.Println("Dato ingresado:",datoEntrada)
	//fmt.Println("Largo dato de entrada:",largoTotal)

	tipo, largoNum, valor, errores := rescataValores(0, datoEntrada)

	if errores != "" {
		return mapa, errores
	}

	//imprimeLog(tipo,largoNum,valor,0)
	salida := validaTipo(tipo, valor)
	if !salida {
		errores = "Valor de dato no corresponde al tipo."
		return mapa, errores
	}

	mapa[tipo[1:3]] = valor

	if (larLargo + largoNum) < largoTotal {
		var totalLargo int = largoNum + larLargo
		for i := 1; i < largoTotal; i++ {

			tipo, largoNum, valor, errores := rescataValores(totalLargo, datoEntrada)

			if errores != "" {
				return mapa, errores
			}

			//imprimeLog(tipo,largoNum,valor,i)

			salida = validaTipo(tipo, valor)

			if !salida {
				errores = "Valor de dato no corresponde al tipo."
				return mapa, errores
			}

			mapa[tipo[1:3]] = valor

			if (totalLargo + larLargo + largoNum) < largoTotal {
				totalLargo = totalLargo + larLargo + largoNum
			} else {
				i = largoTotal
			}
		}
	}

	return mapa, errores
}

func rescataValores(largoAcomulado int, datoEntrada string) (string, int, string, string) {
	var tipo string = ""
	var valor string = ""
	var errores string = ""

	tipo = datoEntrada[largoAcomulado : largoAcomulado+larTipo]

	largoNum, err := strconv.Atoi(datoEntrada[largoAcomulado+larTipo : largoAcomulado+larLargo])

	if err != nil {
		errores = "Error en convertir el largo a numerico: " + datoEntrada[largoAcomulado+larTipo:largoAcomulado+larLargo]
		return tipo, 0, valor, errores
	}

	valor = datoEntrada[largoAcomulado+larLargo : largoAcomulado+larLargo+largoNum]

	return tipo, largoNum, valor, errores

}

func validaTipo(tipo string, valor string) bool {
	tipoDato := tipo[:1]
	if tipoDato != "N" &&
		tipoDato != "n" &&
		tipoDato != "A" &&
		tipoDato != "a" {
		return false
	}

	if tipoDato == "N" || tipoDato == "n" {
		_, err := strconv.Atoi(valor)
		if err == nil {
			return true
		} else {
			return false
		}
	}
	return true
}

func imprimeLog(tipo string, largo int, valor string, vuelta int) {
	fmt.Println("Tipo:", tipo)
	fmt.Println("Largo:", largo)
	fmt.Println("Valor:", valor)
	fmt.Println("Vuelta:", vuelta)
	fmt.Println("======================")
}

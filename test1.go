package main

import (
	"fmt"
)

func main() {
	b := []byte("A0511AB398765UJ1N230200")
	mapa, err := separaDatos(b)
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println(mapa)
	}
}

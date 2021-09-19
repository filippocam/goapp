package main

import (
	"fmt"
	"github.com/filippocam/filonummanip/calc"
	calcNew "github.com/filippocam/filonummanip/v2/calc"
	"log"
)

func main() {
	fmt.Println("sta vedendo Niccolò")
	fmt.Println("MAINNNNNN sto sul nuovo branch")


	result := calc.Add(1)
	fmt.Println("old somma: ", result)
	calc.Echo("ciao", "ciao2")

	err, result := calcNew.Add(1,2,3,4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("new somma: ", result)

	calcNew.Echo("ciao")
}

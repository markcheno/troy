package main

import (
	"github.com/ironbay/troy"
	"log"
)

func main() {
	troy.Update("quigon").Out("taught").V("obiwan").Out("taught").V("anakin").Exec()
	troy.Update("obiwan").Out("taught").V("luke").Exec()

	log.Println(troy.Get("obiwan").Out("taught").Exec())

}

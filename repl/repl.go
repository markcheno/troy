package main

import (
	"github.com/ironbay/troy"
	"log"
)

func main() {
	troy.Update("darth-maul").Out("killed").V("quigon").Out("taught").V("obiwan").Out("taught").V("anakin").Exec()
	troy.Update("obiwan").Out("killed").V("darth-maul").Exec()
	troy.Update("obiwan").Out("taught").V("luke").Exec()
	troy.Update("emperor").Out("taught").V("darth-maul").Exec()

	result := troy.Get("darth-maul").Out("killed").All()
	log.Println("Who did darth-maul kill?", result.Vertices)

	result = troy.Get("obiwan").Out("taught").All()
	log.Println("Who did obiwan teach?", result.Vertices)

	result = troy.Get("darth-maul").Out("killed").V("obiwan")
	log.Println("Did darth-maul kill obiwan?", len(result.Vertices) > 0)
}

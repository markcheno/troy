package main

import (
	"github.com/ironbay/troy"
	"github.com/ironbay/troy/store/cassandra"
	"log"
)

func main() {
	var store cassandra.Store
	store.Create("localhost", "troy")
	troy.Init(&store)

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

package main

import (
	"github.com/ironbay/troy"
	"github.com/ironbay/troy/store"
	"gopkg.in/alecthomas/kingpin.v1"
	"os"
)

var (
	app     = kingpin.New("troy", "Triple store")
	replCmd = app.Command("repl", "Command line interface for debugging")
)

func main() {
	var store store.Memory
	store.Create()
	troy.Init(&store)

	troy.Update("darth-maul").Out("killed").V("quigon").Out("taught").V("obiwan").Out("taught").V("anakin").Exec()
	troy.Update("obiwan").Out("killed").V("darth-maul").Exec()
	troy.Update("obiwan").Out("taught").V("luke").Exec()
	troy.Update("emperor").Out("taught").V("darth-maul").Exec()

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case replCmd.FullCommand():
		repl()
	}

}

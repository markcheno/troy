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

	loadCmd  = app.Command("load", "Command line interface for debugging")
	loadFile = loadCmd.Arg("file", "Input csv file").Required().String()

	webCmd = app.Command("web", "Start web api")
)

func main() {
	var store store.Cassandra
	store.Create("localhost", "troy")
	/*
		var store store.Memory
		store.Create()
	*/
	troy.Init(&store)

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case replCmd.FullCommand():
		repl()
		break
	case loadCmd.FullCommand():
		load(loadFile)
		break
	case webCmd.FullCommand():
		web()
		break
	}

}

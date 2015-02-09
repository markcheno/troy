package main

import (
	"fmt"
	"github.com/ironbay/troy"
	"github.com/peterh/liner"
	"log"
	"os"
	"time"
)

func repl() {

	line := liner.NewLiner()
	defer line.Close()

	history := "/tmp/.troy"
	if f, err := os.Open(history); err == nil {
		line.ReadHistory(f)
		f.Close()
	}

	for {
		l, err := line.Prompt("troy> ")
		if err != nil || l == "exit" || l == "quit" {
			break
		}
		if l == "" {
			continue
		}
		line.AppendHistory(l)
		start := time.Now()
		fmt.Println(troy.Script(l))
		fmt.Println(time.Since(start))
	}

	if f, err := os.Create(history); err != nil {
		log.Print("Error writing history file: ", err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}

}

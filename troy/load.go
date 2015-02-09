package main

import (
	"encoding/csv"
	"github.com/ironbay/troy"
	"io"
	"log"
	"os"
	"sync"
)

func load(path *string) {

	file, err := os.Open(*path)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var wg sync.WaitGroup
	for {
		// read just one record, but we could ReadAll() as well
		record, err := reader.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("Error:", err)
			return
		}

		wg.Add(1)
		troy.Update(record[0]).Out(record[1]).V(record[2]).Exec()
		wg.Done()
	}
	wg.Wait()
}

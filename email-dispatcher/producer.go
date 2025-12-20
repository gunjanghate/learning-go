package main

import (
	"encoding/csv"
	// "fmt"
	"os"
)

func loadRecipient(filePath string, receChan chan Recipient) error {
	defer close(receChan)
	file, err := os.Open(filePath)
	if err != nil{
		return err
	}
	defer file.Close() // run after function end 
	r := csv.NewReader(file)
	rec, err := r.ReadAll()
	if err != nil{
		return err
	}
	
	for _, record := range rec[1:]{
		// fmt.Println(record)
		receChan <- Recipient{
			Name:  record[0],
			Email: record[1],
		}

	}
	return nil
}
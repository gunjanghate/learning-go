package main

import (
	"fmt"
	"html/template"
	"sync"
	"bytes"
)

type Recipient struct{
	Name  string
	Email string
}

// Producer sends email tasks to a queue
// Consumer processes email tasks from the queue
		// send -> consumer with help of channels
		// sender will be blocked until consumer process the data
func main(){
	fmt.Println("Email Dispatcher Service Started")
	receChan := make(chan Recipient) // unbuffered channel


	go func(){
		loadRecipient("./emails.csv", receChan)
	}()

    var wg sync.WaitGroup
	wrkrCnt := 5

	for i:=1; i< wrkrCnt; i++{
		wg.Add(1)
		go emailWorker(i, receChan, &wg)
	}

	wg.Wait()
	fmt.Println("All emails sent successfully")
	

}




func execTmpl(r Recipient) (string, error){
	t, err := template.ParseFiles("./email.tmpl")
	if err != nil{
		return "", err
	}

	var tpl bytes.Buffer

	err1 := t.Execute(&tpl, r)
	if err1 != nil{
		return "", err1
	}

	return tpl.String(), nil

}
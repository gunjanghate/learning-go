package main

import (
	"fmt"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for rec := range ch {
		smtpHost := "localhost"
		smtpPort := 1025

		// formatMsg := fmt.Sprintf("To: %s\r\n Subject: Test Email\r\n\r\n%s\r\n", rec.Email, "Testing")
		// msg := []byte(formatMsg)

		msg, err1 := execTmpl(rec)
		if err1 != nil {
			fmt.Println("Error executing template for ", rec.Email, ":", err1)
			continue
		}
		
		fmt.Println("Worker", id, "sent email to", rec.Email)
		err := smtp.SendMail(smtpHost+":"+fmt.Sprint(smtpPort), nil, "ghategunjan@gmail.com",[]string{rec.Email}, []byte(msg))
		if err != nil {
			fmt.Println("Error sending email to ", rec.Email, ":", err)
			continue
		}
		time.Sleep(50 * time.Millisecond) // Simulate time taken to send email
		fmt.Println("Worker", id, "successfully sent email to", rec.Email)

	}
}
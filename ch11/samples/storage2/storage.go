package storage

import (
	"net/smtp"
	"log"
	"fmt"
)

var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender,
		[]string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendEmail(%s) failed: %s", username, err)
	}
}

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 10000000000 // 1GB
	percent := 100 * used /quota
	if percent < 90 {
		return // OK
	}
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}

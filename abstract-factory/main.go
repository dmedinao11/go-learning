package main

/*
Implementando el patrón de diseño
*/

import (
	"fmt"

	"github.com/dmedinao11/go-learning/abstract-factory/models"
	"github.com/dmedinao11/go-learning/abstract-factory/models/impl"
)

func main() {
	emailNotification := impl.EmailNotification{}
	sendNotification(emailNotification)
	fmt.Println(emailNotification.GetSender().GetSenderMethod())
	fmt.Println(emailNotification.GetSender().GetSenderChannel())
}

func sendNotification(notificationType models.INotificationFactory) {
	notificationType.SendNotification()
}

package impl

import (
	"fmt"

	"github.com/dmedinao11/go-learning/abstract-factory/models"
)

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Enviando email")
}

func (EmailNotification) GetSender() models.ISender {
	return EmailNotificationSender{}
}

package models

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

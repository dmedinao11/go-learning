package models

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

package actor

type EmailSender interface {
	SendEmail() error
}

type SmsSender interface {
	SendSms() error
}

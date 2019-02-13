package port

type EmailSender interface {
	SendEmail() error
}

type SmsSender interface {
	SendSms() error
}

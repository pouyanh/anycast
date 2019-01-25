package infrastructure

type Services struct {
	Logger
	LevelledLogger

	PubSubMessaging
	ReqRepMessaging

	KeyValueStorage
}

package domain

// Message описывает сообщение, которое можно отправить через разные каналы:
// email, SMS, push-уведомление и так далее.
type Message struct {
	To   string
	Text string
}

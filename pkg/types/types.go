package types

type MessageType string

const (
	MessageTypeToday     MessageType = "today"
	MessageTypeInAdvance MessageType = "in_advance"
)

type Message struct {
	Name        string
	Description string
	Date        string
	Type        MessageType
}

type TelegramUser struct {
	UserID string
}

func (u TelegramUser) Recipient() string {
	return u.UserID
}

package protocol

const (
	NewFriendApplication = iota + 1
	FriendApplicationStatusChange
	FriendApplicationNewQuest
	FriendApplicationNewAns
	NewMessage
)

type Notification struct {
	ServerId  int32       `json:"server_id"`
	UserId    uint        `json:"user_id"`
	Type      int8        `json:"type"`
	Content   interface{} `json:"content"`
	Timestamp int64       `json:"timestamp"`
}

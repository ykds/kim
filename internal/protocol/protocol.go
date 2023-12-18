package protocol

const (
	NewFriendApplication          = 1
	FriendApplicationStatusChange = NewFriendApplication << 1
	FriendApplicationNewQuest     = NewFriendApplication << 2
	FriendApplicationNewAns       = NewFriendApplication << 3
)

type Notification struct {
	ServerId  int32       `json:"server_id"`
	UserId    uint        `json:"user_id"`
	Type      int8        `json:"type"`
	Content   interface{} `json:"content"`
	Timestamp int64       `json:"timestamp"`
}

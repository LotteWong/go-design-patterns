package mediator

import (
	"fmt"
	"time"
)

type ChatRoom struct{}

func (c *ChatRoom) ShowMsg(user *User, msg string) {
	fmt.Println(time.Now(), user.GetName(), ":", msg)
}

type User struct {
	name string
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SendMsg(chatRoom *ChatRoom, msg string) {
	chatRoom.ShowMsg(u, msg)
}

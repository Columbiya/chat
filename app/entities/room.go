package entities

type Room struct {
	user1 *User
	user2 *User
}

func (r *Room) Join(u *User) {
}

func (r *Room) Leave(u *User) {

}
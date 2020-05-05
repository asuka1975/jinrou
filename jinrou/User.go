package jinrou

type IUser interface {
	Action(other IUser) string
	canDie() bool
	setProtected()
	AddOnDied(handler func())
	Die()
}

type User struct {
	Name string
	Role string
	protected bool
	onDied func()
}

func (u *User) die() {
	if u.canDie() {
		u.onDied()
	}
}

func (u *User) canDie() bool {
	return !u.protected
}

func (u *User) setProtected() {
	u.protected = true
}

type Villager struct {
	User
}

func newVillager(name string) *Villager {
	v := Villager{ User { Name: name, Role: "Villager", protected: false, onDied: func(){}} }
	return &v
}

func (v *Villager) Action(other IUser) string {
	return ""
}

type Werewolf struct {
	User
}

func newWerewolf(name string) *Werewolf {
	w := Werewolf{ User{ Name: name, Role: "Werewolf", protected: false, onDied: func(){}} }
	return &w
}

func (w *Werewolf) Action(other IUser) string {
	other.Die()
	return ""
}

//TODO add new role

func CreateUser(name string, role string) *IUser {
	switch role {
	case "Werewolf":
		return interface{}(newWerewolf(name)).(*IUser)
	default:
		return interface{}(newVillager(name)).(*IUser)
	}
}
package jinrou

type ISession interface {
	Next() ISession
	PushCommand(queue CommandQueue)
	End()
	Interactive() bool
	GetID() SessionID
}

type SessionID int

const (
	Morning SessionID = 1
	Noon    SessionID = 2
	Evening SessionID = 4
	Night   SessionID = 8
)

type MorningSession struct {
}

func (s MorningSession) Next() ISession {
	return &NoonSession{}
}

func (s MorningSession) PushCommand(queue CommandQueue) {

}

func (s MorningSession) End() {

}

func (s MorningSession) Interactive() bool {
	return false
}

func (s MorningSession) GetID() SessionID {
	return Morning
}

type NoonSession struct {
}

func (s NoonSession) Next() ISession {
	return &EveningSession{}
}

func (s NoonSession) PushCommand(queue CommandQueue) {

}

func (s NoonSession) End() {

}

func (s NoonSession) Interactive() bool {
	return true
}

func (s NoonSession) GetID() SessionID {
	return Noon
}

type EveningSession struct {
}

func (s EveningSession) Next() ISession {
	return &NightSession{}
}

func (s EveningSession) PushCommand(queue CommandQueue) {

}

func (s EveningSession) End() {

}

func (s EveningSession) Interactive() bool {
	return false
}

func (s EveningSession) GetID() SessionID {
	return Evening
}

type NightSession struct {
}

func (s NightSession) Next() ISession {
	return &MorningSession{}
}

func (s NightSession) PushCommand(queue CommandQueue) {

}

func (s NightSession) End() {

}

func (s NightSession) Interactive() bool {
	return false
}

func (s NightSession) GetID() SessionID {
	return Night
}

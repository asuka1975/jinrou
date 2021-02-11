package jinrou

type ISession interface {
	Next() ISession
	String() string
}

type SessionID int

const (
	Morning SessionID = 1
	Noon    SessionID = 2
	Evening SessionID = 4
	Night   SessionID = 8
)

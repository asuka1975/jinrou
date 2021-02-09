package jinrou

type ISession interface {
	Next() ISession
	String() string
	Act(actor string, target string)
	Done()
}

package bankaccount

type Account interface {
	Withdraw(money int) int
	Deposit(money int)
	Balance() int
}

type innerAccount struct {
	balance int
}

func NewAccount() Account {
	return &innerAccount{balance: 0}
}

func (a *innerAccount) Withdraw(money int) int {
	a.balance -= money
	return a.balance
}

func (a *innerAccount) Deposit(money int) {
	a.balance += money
}

func (a *innerAccount) Balance() int {
	return a.balance
}

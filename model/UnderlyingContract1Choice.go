package model

// Choice of the type of underlying contract.
type UnderlyingContract1Choice struct {

	// Underlying registered contract is a loan.
	Loan *LoanContract1 `xml:"Ln"`

	// Underlying registered contract is a commercial trade.
	Trade *TradeContract1 `xml:"Trad"`
}

func (u *UnderlyingContract1Choice) AddLoan() *LoanContract1 {
	u.Loan = new(LoanContract1)
	return u.Loan
}

func (u *UnderlyingContract1Choice) AddTrade() *TradeContract1 {
	u.Trade = new(TradeContract1)
	return u.Trade
}

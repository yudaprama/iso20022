package model

// Balance breakdown information.
type PortfolioBalance1 struct {

	// Summary balance information.
	SummaryBalance []*BalanceDetails5 `xml:"SummryBal"`

	// Detailed balance information.
	DetailedBalance []*BalanceDetails6 `xml:"DtldBal"`
}

func (p *PortfolioBalance1) AddSummaryBalance() *BalanceDetails5 {
	newValue := new(BalanceDetails5)
	p.SummaryBalance = append(p.SummaryBalance, newValue)
	return newValue
}

func (p *PortfolioBalance1) AddDetailedBalance() *BalanceDetails6 {
	newValue := new(BalanceDetails6)
	p.DetailedBalance = append(p.DetailedBalance, newValue)
	return newValue
}

package model

// Choice between the investment account and the financial instrument.
type InvestmentAccountOrFinancialInstrument1Choice struct {

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccount *InvestmentAccount13 `xml:"InvstmtAcct"`

	// Instrument that has intrinsic monetary value, and may transfer value, the price of which may be obtained from a financial market, eg, a bond or a cheque.
	FinancialInstrument *FinancialInstrument6 `xml:"FinInstrm"`
}

func (i *InvestmentAccountOrFinancialInstrument1Choice) AddInvestmentAccount() *InvestmentAccount13 {
	i.InvestmentAccount = new(InvestmentAccount13)
	return i.InvestmentAccount
}

func (i *InvestmentAccountOrFinancialInstrument1Choice) AddFinancialInstrument() *FinancialInstrument6 {
	i.FinancialInstrument = new(FinancialInstrument6)
	return i.FinancialInstrument
}

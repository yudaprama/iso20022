package model

// Overall holding position, in a single financial instrument, held in a securities account at a specified place of safekeeping.
type AggregateHoldingBalance3 struct {

	// Report on the net position of a  financial instrument on the sub-account, for a certain date.
	BalanceForAccount []*AggregateHoldingBalance1 `xml:"BalForAcct"`

	// Agent of the financial instrument, for example, a trade intermediary.
	Agent []*Intermediary29 `xml:"Agt,omitempty"`
}

func (a *AggregateHoldingBalance3) AddBalanceForAccount() *AggregateHoldingBalance1 {
	newValue := new(AggregateHoldingBalance1)
	a.BalanceForAccount = append(a.BalanceForAccount, newValue)
	return newValue
}

func (a *AggregateHoldingBalance3) AddAgent() *Intermediary29 {
	newValue := new(Intermediary29)
	a.Agent = append(a.Agent, newValue)
	return newValue
}

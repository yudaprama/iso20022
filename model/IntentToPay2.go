package model

// Specifies the details of an intention to pay based on purchase orders or commercial invoice.
type IntentToPay2 struct {

	// Specifies if breakdown is by purchase order or commercial invoice.
	Breakdown *BreakDown1Choice `xml:"Brkdwn"`

	// Date at which the payment would be effected.
	ExpectedPaymentDate *ISODate `xml:"XpctdPmtDt"`

	// Specifies the beneficiary's account information.
	SettlementTerms *SettlementTerms3 `xml:"SttlmTerms,omitempty"`
}

func (i *IntentToPay2) AddBreakdown() *BreakDown1Choice {
	i.Breakdown = new(BreakDown1Choice)
	return i.Breakdown
}

func (i *IntentToPay2) SetExpectedPaymentDate(value string) {
	i.ExpectedPaymentDate = (*ISODate)(&value)
}

func (i *IntentToPay2) AddSettlementTerms() *SettlementTerms3 {
	i.SettlementTerms = new(SettlementTerms3)
	return i.SettlementTerms
}

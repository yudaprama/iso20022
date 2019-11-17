package model

// Additional specific query criteria.
type AdditionalQueryParameters11 struct {

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status.
	Status *Status19Choice `xml:"Sts,omitempty"`

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status reason.
	Reason []*Reason16Choice `xml:"Rsn,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification []*SecurityIdentification19 `xml:"FinInstrmId,omitempty"`
}

func (a *AdditionalQueryParameters11) AddStatus() *Status19Choice {
	a.Status = new(Status19Choice)
	return a.Status
}

func (a *AdditionalQueryParameters11) AddReason() *Reason16Choice {
	newValue := new(Reason16Choice)
	a.Reason = append(a.Reason, newValue)
	return newValue
}

func (a *AdditionalQueryParameters11) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	newValue := new(SecurityIdentification19)
	a.FinancialInstrumentIdentification = append(a.FinancialInstrumentIdentification, newValue)
	return newValue
}

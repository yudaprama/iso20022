package model

// Additional specific query criteria.
type AdditionalQueryParameters7 struct {

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status.
	Status *Status8Choice `xml:"Sts,omitempty"`

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status reason.
	Reason []*Reason12Choice `xml:"Rsn,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification []*SecurityIdentification14 `xml:"FinInstrmId,omitempty"`
}

func (a *AdditionalQueryParameters7) AddStatus() *Status8Choice {
	a.Status = new(Status8Choice)
	return a.Status
}

func (a *AdditionalQueryParameters7) AddReason() *Reason12Choice {
	newValue := new(Reason12Choice)
	a.Reason = append(a.Reason, newValue)
	return newValue
}

func (a *AdditionalQueryParameters7) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	newValue := new(SecurityIdentification14)
	a.FinancialInstrumentIdentification = append(a.FinancialInstrumentIdentification, newValue)
	return newValue
}

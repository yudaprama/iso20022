package model

// Additional specific query criteria.
type AdditionalQueryParameters3 struct {

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status.
	Status *Status8Choice `xml:"Sts,omitempty"`

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status reason.
	Reason []*Reason6Choice `xml:"Rsn,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification []*SecurityIdentification14 `xml:"FinInstrmId,omitempty"`
}

func (a *AdditionalQueryParameters3) AddStatus() *Status8Choice {
	a.Status = new(Status8Choice)
	return a.Status
}

func (a *AdditionalQueryParameters3) AddReason() *Reason6Choice {
	newValue := new(Reason6Choice)
	a.Reason = append(a.Reason, newValue)
	return newValue
}

func (a *AdditionalQueryParameters3) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	newValue := new(SecurityIdentification14)
	a.FinancialInstrumentIdentification = append(a.FinancialInstrumentIdentification, newValue)
	return newValue
}

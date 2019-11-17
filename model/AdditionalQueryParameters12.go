package model

// Additional specific query criteria.
type AdditionalQueryParameters12 struct {

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status.
	Status *Status22Choice `xml:"Sts,omitempty"`

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status reason.
	Reason []*Reason17Choice `xml:"Rsn,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification []*SecurityIdentification20 `xml:"FinInstrmId,omitempty"`
}

func (a *AdditionalQueryParameters12) AddStatus() *Status22Choice {
	a.Status = new(Status22Choice)
	return a.Status
}

func (a *AdditionalQueryParameters12) AddReason() *Reason17Choice {
	newValue := new(Reason17Choice)
	a.Reason = append(a.Reason, newValue)
	return newValue
}

func (a *AdditionalQueryParameters12) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	newValue := new(SecurityIdentification20)
	a.FinancialInstrumentIdentification = append(a.FinancialInstrumentIdentification, newValue)
	return newValue
}

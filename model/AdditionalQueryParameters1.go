package model

// Additional specific query criteria.
type AdditionalQueryParameters1 struct {

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status.
	Status *Status1Choice `xml:"Sts,omitempty"`

	// Request to obtain a Securities Transaction Pending Report for transactions with the specified status reason.
	Reason []*Reason1Choice `xml:"Rsn,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification []*SecurityIdentification11 `xml:"FinInstrmId,omitempty"`
}

func (a *AdditionalQueryParameters1) AddStatus() *Status1Choice {
	a.Status = new(Status1Choice)
	return a.Status
}

func (a *AdditionalQueryParameters1) AddReason() *Reason1Choice {
	newValue := new(Reason1Choice)
	a.Reason = append(a.Reason, newValue)
	return newValue
}

func (a *AdditionalQueryParameters1) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	newValue := new(SecurityIdentification11)
	a.FinancialInstrumentIdentification = append(a.FinancialInstrumentIdentification, newValue)
	return newValue
}

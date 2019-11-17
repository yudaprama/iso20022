package model

// Unique identifier of an account, as assigned by the account servicer.
type AccountIdentification1 struct {

	// Unique identifier for an account. It is assigned by the account servicer using a proprietary identification scheme.
	Proprietary *SimpleIdentificationInformation `xml:"Prtry"`
}

func (a *AccountIdentification1) AddProprietary() *SimpleIdentificationInformation {
	a.Proprietary = new(SimpleIdentificationInformation)
	return a.Proprietary
}

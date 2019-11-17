package model

// Unique identifier of an account, as assigned by the account servicer.
type AccountIdentification4 struct {

	// Unique identifier for an account. It is assigned by the account servicer using a proprietary identification scheme.
	Proprietary *SimpleIdentificationInformation1 `xml:"Prtry"`
}

func (a *AccountIdentification4) AddProprietary() *SimpleIdentificationInformation1 {
	a.Proprietary = new(SimpleIdentificationInformation1)
	return a.Proprietary
}

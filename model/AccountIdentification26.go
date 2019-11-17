package model

// Unique identifier of an account, as assigned by the account servicer.
type AccountIdentification26 struct {

	// Unique identifier for an account. It is assigned by the account servicer using a proprietary identification scheme.
	Proprietary *SimpleIdentificationInformation4 `xml:"Prtry"`
}

func (a *AccountIdentification26) AddProprietary() *SimpleIdentificationInformation4 {
	a.Proprietary = new(SimpleIdentificationInformation4)
	return a.Proprietary
}

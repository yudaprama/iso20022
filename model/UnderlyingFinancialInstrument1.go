package model

// Underlying financial instrument to which an trade confirmation is related.
type UnderlyingFinancialInstrument1 struct {

	// Identification of the underlying security by an ISIN.
	Identification *SecurityIdentification14 `xml:"Id"`

	// Underlying financial instrument attributes to which an trade confirmation is related.
	Attributes *FinancialInstrumentAttributes31 `xml:"Attrbts,omitempty"`
}

func (u *UnderlyingFinancialInstrument1) AddIdentification() *SecurityIdentification14 {
	u.Identification = new(SecurityIdentification14)
	return u.Identification
}

func (u *UnderlyingFinancialInstrument1) AddAttributes() *FinancialInstrumentAttributes31 {
	u.Attributes = new(FinancialInstrumentAttributes31)
	return u.Attributes
}

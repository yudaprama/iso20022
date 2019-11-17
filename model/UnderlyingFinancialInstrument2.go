package model

// Underlying financial instrument to which an trade confirmation is related.
type UnderlyingFinancialInstrument2 struct {

	// Identification of the underlying security by an ISIN.
	Identification *SecurityIdentification14 `xml:"Id"`

	// Underlying financial instrument attributes to which an trade confirmation is related.
	Attributes *FinancialInstrumentAttributes44 `xml:"Attrbts,omitempty"`
}

func (u *UnderlyingFinancialInstrument2) AddIdentification() *SecurityIdentification14 {
	u.Identification = new(SecurityIdentification14)
	return u.Identification
}

func (u *UnderlyingFinancialInstrument2) AddAttributes() *FinancialInstrumentAttributes44 {
	u.Attributes = new(FinancialInstrumentAttributes44)
	return u.Attributes
}

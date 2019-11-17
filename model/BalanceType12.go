package model

// Set of elements used to define the balance type and sub-type.
type BalanceType12 struct {

	// Coded or proprietary format balance type.
	CodeOrProprietary *BalanceType5Choice `xml:"CdOrPrtry"`

	// Specifies the balance sub-type.
	SubType *BalanceSubType1Choice `xml:"SubTp,omitempty"`
}

func (b *BalanceType12) AddCodeOrProprietary() *BalanceType5Choice {
	b.CodeOrProprietary = new(BalanceType5Choice)
	return b.CodeOrProprietary
}

func (b *BalanceType12) AddSubType() *BalanceSubType1Choice {
	b.SubType = new(BalanceSubType1Choice)
	return b.SubType
}

package model

// Choice of formats for the tax basis.
type TaxBasis1Choice struct {

	// Tax basis expressed as a code.
	Code *TaxationBasis2Code `xml:"Cd"`

	// Tax basis expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TaxBasis1Choice) SetCode(value string) {
	t.Code = (*TaxationBasis2Code)(&value)
}

func (t *TaxBasis1Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}

package model

// Choice between an indicator or a proprietary code to specify whether the security is a temporary security.
type TemporaryFinancialInstrumentIndicator1Choice struct {

	// Temporary financial instrument identification used for processing reasons.
	TemporaryIndicator *YesNoIndicator `xml:"TempInd"`

	// Proprietary code to specify whether the security is a temporary security.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (t *TemporaryFinancialInstrumentIndicator1Choice) SetTemporaryIndicator(value string) {
	t.TemporaryIndicator = (*YesNoIndicator)(&value)
}

func (t *TemporaryFinancialInstrumentIndicator1Choice) AddProprietary() *GenericIdentification20 {
	t.Proprietary = new(GenericIdentification20)
	return t.Proprietary
}

package model

// Choice between an indicator or a proprietary code to specify whether the security is a temporary security.
type TemporaryFinancialInstrumentIndicator3Choice struct {

	// Temporary financial instrument identification used for processing reasons.
	TemporaryIndicator *YesNoIndicator `xml:"TempInd"`

	// Proprietary code to specify whether the security is a temporary security.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (t *TemporaryFinancialInstrumentIndicator3Choice) SetTemporaryIndicator(value string) {
	t.TemporaryIndicator = (*YesNoIndicator)(&value)
}

func (t *TemporaryFinancialInstrumentIndicator3Choice) AddProprietary() *GenericIdentification30 {
	t.Proprietary = new(GenericIdentification30)
	return t.Proprietary
}

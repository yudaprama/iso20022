package model

// Choice between an indicator or a proprietary code to specify whether the security is a temporary security.
type TemporaryFinancialInstrumentIndicator4Choice struct {

	// Temporary financial instrument identification used for processing reasons.
	TemporaryIndicator *YesNoIndicator `xml:"TempInd"`

	// Proprietary code to specify whether the security is a temporary security.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TemporaryFinancialInstrumentIndicator4Choice) SetTemporaryIndicator(value string) {
	t.TemporaryIndicator = (*YesNoIndicator)(&value)
}

func (t *TemporaryFinancialInstrumentIndicator4Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
